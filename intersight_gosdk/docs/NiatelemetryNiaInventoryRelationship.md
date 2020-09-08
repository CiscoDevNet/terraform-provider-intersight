# NiatelemetryNiaInventoryRelationship

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
**Cpu** | Pointer to **float32** | CPU usage of device being inventoried. This determines the percentage of CPU resources used. | [optional] 
**CrashResetLogs** | Pointer to **string** | Last crash reset reason of device being inventoried. This determines the last reason for a device&#39;s restart due to crash of the system. | [optional] 
**DeviceName** | Pointer to **string** | Name of device being inventoried. The name the user assigns to the device is inventoried here. | [optional] 
**DeviceType** | Pointer to **string** | Type of device being inventoried. This determines whether the device is a controller, leaf or spine. | [optional] 
**Disk** | Pointer to [**NiatelemetryDiskinfo**](niatelemetry.Diskinfo.md) |  | [optional] 
**FexCount** | Pointer to **int64** | Scale of fabric extendors utilized. | [optional] 
**LogInTime** | Pointer to **string** | Last log in time device being inventoried. This determines the last login time on the device. | [optional] 
**LogOutTime** | Pointer to **string** | Last log out time of device being inventoried. This determines the last logout time on the device. | [optional] 
**MacSecCount** | Pointer to **int64** | Number of Macsec configured interfaces on a TOR. | [optional] 
**MacSecFabCount** | Pointer to **int64** | Number of Macsec configured interfaces on a Spine. | [optional] 
**MacsecTotalCount** | Pointer to **int64** | Number of total Macsec configured interfaces for all nodes. | [optional] 
**Memory** | Pointer to **int64** | Memory usage of device being inventoried. This determines the percentage of memory resources used. | [optional] 
**RecordType** | Pointer to **string** | Type of record DCNM / APIC / SE. This determines the type of platform where inventory was collected. | [optional] 
**RecordVersion** | Pointer to **string** | Version of record being pushed. This determines what was the API version for data available from the device. | [optional] 
**RoutePrefixCount** | Pointer to **int64** | Total nuumber of v4 and v6 routes per node. | [optional] 
**RoutePrefixV4Count** | Pointer to **int64** | Scale of v4 routes per node. | [optional] 
**RoutePrefixV6Count** | Pointer to **int64** | Scale of v6 routes per node. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being invetoried. The serial number is unique per device and will be used as the key. | [optional] 
**SoftwareDownload** | Pointer to **string** | Last software downloaded of device being inventoried. This determines if software download API was used. | [optional] 
**Version** | Pointer to **string** | Software version of device being inventoried. The various software version values for each device are available on cisco.com. | [optional] 
**LicenseState** | Pointer to [**NiatelemetryNiaLicenseStateRelationship**](niatelemetry.NiaLicenseState.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaInventoryRelationship

`func NewNiatelemetryNiaInventoryRelationship(classId string, objectType string, ) *NiatelemetryNiaInventoryRelationship`

NewNiatelemetryNiaInventoryRelationship instantiates a new NiatelemetryNiaInventoryRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaInventoryRelationshipWithDefaults

`func NewNiatelemetryNiaInventoryRelationshipWithDefaults() *NiatelemetryNiaInventoryRelationship`

NewNiatelemetryNiaInventoryRelationshipWithDefaults instantiates a new NiatelemetryNiaInventoryRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *NiatelemetryNiaInventoryRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *NiatelemetryNiaInventoryRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *NiatelemetryNiaInventoryRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *NiatelemetryNiaInventoryRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *NiatelemetryNiaInventoryRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaInventoryRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaInventoryRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *NiatelemetryNiaInventoryRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NiatelemetryNiaInventoryRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NiatelemetryNiaInventoryRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *NiatelemetryNiaInventoryRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *NiatelemetryNiaInventoryRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *NiatelemetryNiaInventoryRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *NiatelemetryNiaInventoryRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *NiatelemetryNiaInventoryRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *NiatelemetryNiaInventoryRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *NiatelemetryNiaInventoryRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *NiatelemetryNiaInventoryRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *NiatelemetryNiaInventoryRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *NiatelemetryNiaInventoryRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *NiatelemetryNiaInventoryRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *NiatelemetryNiaInventoryRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaInventoryRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *NiatelemetryNiaInventoryRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *NiatelemetryNiaInventoryRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *NiatelemetryNiaInventoryRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *NiatelemetryNiaInventoryRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *NiatelemetryNiaInventoryRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *NiatelemetryNiaInventoryRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *NiatelemetryNiaInventoryRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *NiatelemetryNiaInventoryRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NiatelemetryNiaInventoryRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NiatelemetryNiaInventoryRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NiatelemetryNiaInventoryRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *NiatelemetryNiaInventoryRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *NiatelemetryNiaInventoryRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *NiatelemetryNiaInventoryRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *NiatelemetryNiaInventoryRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *NiatelemetryNiaInventoryRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *NiatelemetryNiaInventoryRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *NiatelemetryNiaInventoryRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *NiatelemetryNiaInventoryRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *NiatelemetryNiaInventoryRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *NiatelemetryNiaInventoryRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *NiatelemetryNiaInventoryRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *NiatelemetryNiaInventoryRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *NiatelemetryNiaInventoryRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *NiatelemetryNiaInventoryRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *NiatelemetryNiaInventoryRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *NiatelemetryNiaInventoryRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *NiatelemetryNiaInventoryRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *NiatelemetryNiaInventoryRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *NiatelemetryNiaInventoryRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *NiatelemetryNiaInventoryRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *NiatelemetryNiaInventoryRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *NiatelemetryNiaInventoryRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *NiatelemetryNiaInventoryRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *NiatelemetryNiaInventoryRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *NiatelemetryNiaInventoryRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *NiatelemetryNiaInventoryRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetCpu

`func (o *NiatelemetryNiaInventoryRelationship) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *NiatelemetryNiaInventoryRelationship) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *NiatelemetryNiaInventoryRelationship) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *NiatelemetryNiaInventoryRelationship) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCrashResetLogs

`func (o *NiatelemetryNiaInventoryRelationship) GetCrashResetLogs() string`

GetCrashResetLogs returns the CrashResetLogs field if non-nil, zero value otherwise.

### GetCrashResetLogsOk

`func (o *NiatelemetryNiaInventoryRelationship) GetCrashResetLogsOk() (*string, bool)`

GetCrashResetLogsOk returns a tuple with the CrashResetLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrashResetLogs

`func (o *NiatelemetryNiaInventoryRelationship) SetCrashResetLogs(v string)`

SetCrashResetLogs sets CrashResetLogs field to given value.

### HasCrashResetLogs

`func (o *NiatelemetryNiaInventoryRelationship) HasCrashResetLogs() bool`

HasCrashResetLogs returns a boolean if a field has been set.

### GetDeviceName

`func (o *NiatelemetryNiaInventoryRelationship) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *NiatelemetryNiaInventoryRelationship) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *NiatelemetryNiaInventoryRelationship) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *NiatelemetryNiaInventoryRelationship) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *NiatelemetryNiaInventoryRelationship) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *NiatelemetryNiaInventoryRelationship) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *NiatelemetryNiaInventoryRelationship) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDisk

`func (o *NiatelemetryNiaInventoryRelationship) GetDisk() NiatelemetryDiskinfo`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *NiatelemetryNiaInventoryRelationship) GetDiskOk() (*NiatelemetryDiskinfo, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *NiatelemetryNiaInventoryRelationship) SetDisk(v NiatelemetryDiskinfo)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *NiatelemetryNiaInventoryRelationship) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetFexCount

`func (o *NiatelemetryNiaInventoryRelationship) GetFexCount() int64`

GetFexCount returns the FexCount field if non-nil, zero value otherwise.

### GetFexCountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetFexCountOk() (*int64, bool)`

GetFexCountOk returns a tuple with the FexCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFexCount

`func (o *NiatelemetryNiaInventoryRelationship) SetFexCount(v int64)`

SetFexCount sets FexCount field to given value.

### HasFexCount

`func (o *NiatelemetryNiaInventoryRelationship) HasFexCount() bool`

HasFexCount returns a boolean if a field has been set.

### GetLogInTime

`func (o *NiatelemetryNiaInventoryRelationship) GetLogInTime() string`

GetLogInTime returns the LogInTime field if non-nil, zero value otherwise.

### GetLogInTimeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetLogInTimeOk() (*string, bool)`

GetLogInTimeOk returns a tuple with the LogInTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogInTime

`func (o *NiatelemetryNiaInventoryRelationship) SetLogInTime(v string)`

SetLogInTime sets LogInTime field to given value.

### HasLogInTime

`func (o *NiatelemetryNiaInventoryRelationship) HasLogInTime() bool`

HasLogInTime returns a boolean if a field has been set.

### GetLogOutTime

`func (o *NiatelemetryNiaInventoryRelationship) GetLogOutTime() string`

GetLogOutTime returns the LogOutTime field if non-nil, zero value otherwise.

### GetLogOutTimeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetLogOutTimeOk() (*string, bool)`

GetLogOutTimeOk returns a tuple with the LogOutTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogOutTime

`func (o *NiatelemetryNiaInventoryRelationship) SetLogOutTime(v string)`

SetLogOutTime sets LogOutTime field to given value.

### HasLogOutTime

`func (o *NiatelemetryNiaInventoryRelationship) HasLogOutTime() bool`

HasLogOutTime returns a boolean if a field has been set.

### GetMacSecCount

`func (o *NiatelemetryNiaInventoryRelationship) GetMacSecCount() int64`

GetMacSecCount returns the MacSecCount field if non-nil, zero value otherwise.

### GetMacSecCountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetMacSecCountOk() (*int64, bool)`

GetMacSecCountOk returns a tuple with the MacSecCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecCount

`func (o *NiatelemetryNiaInventoryRelationship) SetMacSecCount(v int64)`

SetMacSecCount sets MacSecCount field to given value.

### HasMacSecCount

`func (o *NiatelemetryNiaInventoryRelationship) HasMacSecCount() bool`

HasMacSecCount returns a boolean if a field has been set.

### GetMacSecFabCount

`func (o *NiatelemetryNiaInventoryRelationship) GetMacSecFabCount() int64`

GetMacSecFabCount returns the MacSecFabCount field if non-nil, zero value otherwise.

### GetMacSecFabCountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetMacSecFabCountOk() (*int64, bool)`

GetMacSecFabCountOk returns a tuple with the MacSecFabCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacSecFabCount

`func (o *NiatelemetryNiaInventoryRelationship) SetMacSecFabCount(v int64)`

SetMacSecFabCount sets MacSecFabCount field to given value.

### HasMacSecFabCount

`func (o *NiatelemetryNiaInventoryRelationship) HasMacSecFabCount() bool`

HasMacSecFabCount returns a boolean if a field has been set.

### GetMacsecTotalCount

`func (o *NiatelemetryNiaInventoryRelationship) GetMacsecTotalCount() int64`

GetMacsecTotalCount returns the MacsecTotalCount field if non-nil, zero value otherwise.

### GetMacsecTotalCountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetMacsecTotalCountOk() (*int64, bool)`

GetMacsecTotalCountOk returns a tuple with the MacsecTotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacsecTotalCount

`func (o *NiatelemetryNiaInventoryRelationship) SetMacsecTotalCount(v int64)`

SetMacsecTotalCount sets MacsecTotalCount field to given value.

### HasMacsecTotalCount

`func (o *NiatelemetryNiaInventoryRelationship) HasMacsecTotalCount() bool`

HasMacsecTotalCount returns a boolean if a field has been set.

### GetMemory

`func (o *NiatelemetryNiaInventoryRelationship) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *NiatelemetryNiaInventoryRelationship) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *NiatelemetryNiaInventoryRelationship) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *NiatelemetryNiaInventoryRelationship) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetRecordType

`func (o *NiatelemetryNiaInventoryRelationship) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NiatelemetryNiaInventoryRelationship) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NiatelemetryNiaInventoryRelationship) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRecordVersion

`func (o *NiatelemetryNiaInventoryRelationship) GetRecordVersion() string`

GetRecordVersion returns the RecordVersion field if non-nil, zero value otherwise.

### GetRecordVersionOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRecordVersionOk() (*string, bool)`

GetRecordVersionOk returns a tuple with the RecordVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordVersion

`func (o *NiatelemetryNiaInventoryRelationship) SetRecordVersion(v string)`

SetRecordVersion sets RecordVersion field to given value.

### HasRecordVersion

`func (o *NiatelemetryNiaInventoryRelationship) HasRecordVersion() bool`

HasRecordVersion returns a boolean if a field has been set.

### GetRoutePrefixCount

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixCount() int64`

GetRoutePrefixCount returns the RoutePrefixCount field if non-nil, zero value otherwise.

### GetRoutePrefixCountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixCountOk() (*int64, bool)`

GetRoutePrefixCountOk returns a tuple with the RoutePrefixCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixCount

`func (o *NiatelemetryNiaInventoryRelationship) SetRoutePrefixCount(v int64)`

SetRoutePrefixCount sets RoutePrefixCount field to given value.

### HasRoutePrefixCount

`func (o *NiatelemetryNiaInventoryRelationship) HasRoutePrefixCount() bool`

HasRoutePrefixCount returns a boolean if a field has been set.

### GetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixV4Count() int64`

GetRoutePrefixV4Count returns the RoutePrefixV4Count field if non-nil, zero value otherwise.

### GetRoutePrefixV4CountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixV4CountOk() (*int64, bool)`

GetRoutePrefixV4CountOk returns a tuple with the RoutePrefixV4Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryRelationship) SetRoutePrefixV4Count(v int64)`

SetRoutePrefixV4Count sets RoutePrefixV4Count field to given value.

### HasRoutePrefixV4Count

`func (o *NiatelemetryNiaInventoryRelationship) HasRoutePrefixV4Count() bool`

HasRoutePrefixV4Count returns a boolean if a field has been set.

### GetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixV6Count() int64`

GetRoutePrefixV6Count returns the RoutePrefixV6Count field if non-nil, zero value otherwise.

### GetRoutePrefixV6CountOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRoutePrefixV6CountOk() (*int64, bool)`

GetRoutePrefixV6CountOk returns a tuple with the RoutePrefixV6Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryRelationship) SetRoutePrefixV6Count(v int64)`

SetRoutePrefixV6Count sets RoutePrefixV6Count field to given value.

### HasRoutePrefixV6Count

`func (o *NiatelemetryNiaInventoryRelationship) HasRoutePrefixV6Count() bool`

HasRoutePrefixV6Count returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaInventoryRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaInventoryRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaInventoryRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaInventoryRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetSoftwareDownload

`func (o *NiatelemetryNiaInventoryRelationship) GetSoftwareDownload() string`

GetSoftwareDownload returns the SoftwareDownload field if non-nil, zero value otherwise.

### GetSoftwareDownloadOk

`func (o *NiatelemetryNiaInventoryRelationship) GetSoftwareDownloadOk() (*string, bool)`

GetSoftwareDownloadOk returns a tuple with the SoftwareDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareDownload

`func (o *NiatelemetryNiaInventoryRelationship) SetSoftwareDownload(v string)`

SetSoftwareDownload sets SoftwareDownload field to given value.

### HasSoftwareDownload

`func (o *NiatelemetryNiaInventoryRelationship) HasSoftwareDownload() bool`

HasSoftwareDownload returns a boolean if a field has been set.

### GetVersion

`func (o *NiatelemetryNiaInventoryRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *NiatelemetryNiaInventoryRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *NiatelemetryNiaInventoryRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *NiatelemetryNiaInventoryRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLicenseState

`func (o *NiatelemetryNiaInventoryRelationship) GetLicenseState() NiatelemetryNiaLicenseStateRelationship`

GetLicenseState returns the LicenseState field if non-nil, zero value otherwise.

### GetLicenseStateOk

`func (o *NiatelemetryNiaInventoryRelationship) GetLicenseStateOk() (*NiatelemetryNiaLicenseStateRelationship, bool)`

GetLicenseStateOk returns a tuple with the LicenseState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseState

`func (o *NiatelemetryNiaInventoryRelationship) SetLicenseState(v NiatelemetryNiaLicenseStateRelationship)`

SetLicenseState sets LicenseState field to given value.

### HasLicenseState

`func (o *NiatelemetryNiaInventoryRelationship) HasLicenseState() bool`

HasLicenseState returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *NiatelemetryNiaInventoryRelationship) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *NiatelemetryNiaInventoryRelationship) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *NiatelemetryNiaInventoryRelationship) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *NiatelemetryNiaInventoryRelationship) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


