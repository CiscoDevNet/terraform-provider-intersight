# NiatelemetryNiaLicenseStateRelationship

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
**FeatureActivated** | Pointer to **string** | Features activated on device being inventoried. This determines which features are currently enabled on the device that the license API can check. | [optional] 
**LicenseActivated** | Pointer to **string** | Licenses activated on device being inventoried. This determines which lienceses are currently enabled on the device. | [optional] 
**PidType** | Pointer to **string** | PID of device being inventoried. This determines the hardware model type of the device. | [optional] 
**Serial** | Pointer to **string** | Serial number of device being inventoried. The serial number is unique per device. | [optional] 
**Device** | Pointer to [**NiatelemetryNiaInventoryRelationship**](niatelemetry.NiaInventory.Relationship.md) |  | [optional] 

## Methods

### NewNiatelemetryNiaLicenseStateRelationship

`func NewNiatelemetryNiaLicenseStateRelationship(classId string, objectType string, ) *NiatelemetryNiaLicenseStateRelationship`

NewNiatelemetryNiaLicenseStateRelationship instantiates a new NiatelemetryNiaLicenseStateRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNiatelemetryNiaLicenseStateRelationshipWithDefaults

`func NewNiatelemetryNiaLicenseStateRelationshipWithDefaults() *NiatelemetryNiaLicenseStateRelationship`

NewNiatelemetryNiaLicenseStateRelationshipWithDefaults instantiates a new NiatelemetryNiaLicenseStateRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *NiatelemetryNiaLicenseStateRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *NiatelemetryNiaLicenseStateRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *NiatelemetryNiaLicenseStateRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *NiatelemetryNiaLicenseStateRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *NiatelemetryNiaLicenseStateRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *NiatelemetryNiaLicenseStateRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *NiatelemetryNiaLicenseStateRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *NiatelemetryNiaLicenseStateRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *NiatelemetryNiaLicenseStateRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *NiatelemetryNiaLicenseStateRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *NiatelemetryNiaLicenseStateRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *NiatelemetryNiaLicenseStateRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *NiatelemetryNiaLicenseStateRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *NiatelemetryNiaLicenseStateRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *NiatelemetryNiaLicenseStateRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *NiatelemetryNiaLicenseStateRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *NiatelemetryNiaLicenseStateRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *NiatelemetryNiaLicenseStateRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NiatelemetryNiaLicenseStateRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NiatelemetryNiaLicenseStateRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *NiatelemetryNiaLicenseStateRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *NiatelemetryNiaLicenseStateRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *NiatelemetryNiaLicenseStateRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *NiatelemetryNiaLicenseStateRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *NiatelemetryNiaLicenseStateRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *NiatelemetryNiaLicenseStateRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *NiatelemetryNiaLicenseStateRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *NiatelemetryNiaLicenseStateRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *NiatelemetryNiaLicenseStateRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *NiatelemetryNiaLicenseStateRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *NiatelemetryNiaLicenseStateRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *NiatelemetryNiaLicenseStateRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *NiatelemetryNiaLicenseStateRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *NiatelemetryNiaLicenseStateRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *NiatelemetryNiaLicenseStateRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *NiatelemetryNiaLicenseStateRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *NiatelemetryNiaLicenseStateRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *NiatelemetryNiaLicenseStateRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *NiatelemetryNiaLicenseStateRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *NiatelemetryNiaLicenseStateRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetFeatureActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) GetFeatureActivated() string`

GetFeatureActivated returns the FeatureActivated field if non-nil, zero value otherwise.

### GetFeatureActivatedOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetFeatureActivatedOk() (*string, bool)`

GetFeatureActivatedOk returns a tuple with the FeatureActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) SetFeatureActivated(v string)`

SetFeatureActivated sets FeatureActivated field to given value.

### HasFeatureActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) HasFeatureActivated() bool`

HasFeatureActivated returns a boolean if a field has been set.

### GetLicenseActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) GetLicenseActivated() string`

GetLicenseActivated returns the LicenseActivated field if non-nil, zero value otherwise.

### GetLicenseActivatedOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetLicenseActivatedOk() (*string, bool)`

GetLicenseActivatedOk returns a tuple with the LicenseActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) SetLicenseActivated(v string)`

SetLicenseActivated sets LicenseActivated field to given value.

### HasLicenseActivated

`func (o *NiatelemetryNiaLicenseStateRelationship) HasLicenseActivated() bool`

HasLicenseActivated returns a boolean if a field has been set.

### GetPidType

`func (o *NiatelemetryNiaLicenseStateRelationship) GetPidType() string`

GetPidType returns the PidType field if non-nil, zero value otherwise.

### GetPidTypeOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetPidTypeOk() (*string, bool)`

GetPidTypeOk returns a tuple with the PidType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPidType

`func (o *NiatelemetryNiaLicenseStateRelationship) SetPidType(v string)`

SetPidType sets PidType field to given value.

### HasPidType

`func (o *NiatelemetryNiaLicenseStateRelationship) HasPidType() bool`

HasPidType returns a boolean if a field has been set.

### GetSerial

`func (o *NiatelemetryNiaLicenseStateRelationship) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *NiatelemetryNiaLicenseStateRelationship) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *NiatelemetryNiaLicenseStateRelationship) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDevice

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDevice() NiatelemetryNiaInventoryRelationship`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *NiatelemetryNiaLicenseStateRelationship) GetDeviceOk() (*NiatelemetryNiaInventoryRelationship, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *NiatelemetryNiaLicenseStateRelationship) SetDevice(v NiatelemetryNiaInventoryRelationship)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *NiatelemetryNiaLicenseStateRelationship) HasDevice() bool`

HasDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


