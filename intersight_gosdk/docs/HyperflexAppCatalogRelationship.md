# HyperflexAppCatalogRelationship

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
**Version** | Pointer to **string** | The catalog version used in HyperFlex cluster configuration service. | [optional] 
**FeatureLimitExternal** | Pointer to [**HyperflexFeatureLimitExternalRelationship**](hyperflex.FeatureLimitExternal.Relationship.md) |  | [optional] 
**FeatureLimitInternal** | Pointer to [**HyperflexFeatureLimitInternalRelationship**](hyperflex.FeatureLimitInternal.Relationship.md) |  | [optional] 
**HxdpVersions** | Pointer to [**[]HyperflexHxdpVersionRelationship**](hyperflex.HxdpVersion.Relationship.md) | An array of relationships to hyperflexHxdpVersion resources. | [optional] 
**HyperflexCapabilityInfos** | Pointer to [**[]HyperflexCapabilityInfoRelationship**](hyperflex.CapabilityInfo.Relationship.md) | An array of relationships to hyperflexCapabilityInfo resources. | [optional] 
**HyperflexSoftwareCompatibilityInfos** | Pointer to [**[]HclHyperflexSoftwareCompatibilityInfoRelationship**](hcl.HyperflexSoftwareCompatibilityInfo.Relationship.md) | An array of relationships to hclHyperflexSoftwareCompatibilityInfo resources. | [optional] 
**ServerFirmwareVersion** | Pointer to [**HyperflexServerFirmwareVersionRelationship**](hyperflex.ServerFirmwareVersion.Relationship.md) |  | [optional] 
**ServerModel** | Pointer to [**HyperflexServerModelRelationship**](hyperflex.ServerModel.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexAppCatalogRelationship

`func NewHyperflexAppCatalogRelationship(classId string, objectType string, ) *HyperflexAppCatalogRelationship`

NewHyperflexAppCatalogRelationship instantiates a new HyperflexAppCatalogRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexAppCatalogRelationshipWithDefaults

`func NewHyperflexAppCatalogRelationshipWithDefaults() *HyperflexAppCatalogRelationship`

NewHyperflexAppCatalogRelationshipWithDefaults instantiates a new HyperflexAppCatalogRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexAppCatalogRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexAppCatalogRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexAppCatalogRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexAppCatalogRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexAppCatalogRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexAppCatalogRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexAppCatalogRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexAppCatalogRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexAppCatalogRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexAppCatalogRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexAppCatalogRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexAppCatalogRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexAppCatalogRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexAppCatalogRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexAppCatalogRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexAppCatalogRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexAppCatalogRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexAppCatalogRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexAppCatalogRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexAppCatalogRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexAppCatalogRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexAppCatalogRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexAppCatalogRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexAppCatalogRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexAppCatalogRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexAppCatalogRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexAppCatalogRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexAppCatalogRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexAppCatalogRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexAppCatalogRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexAppCatalogRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexAppCatalogRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexAppCatalogRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexAppCatalogRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexAppCatalogRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexAppCatalogRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexAppCatalogRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexAppCatalogRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexAppCatalogRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexAppCatalogRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexAppCatalogRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexAppCatalogRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexAppCatalogRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexAppCatalogRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexAppCatalogRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexAppCatalogRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexAppCatalogRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexAppCatalogRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexAppCatalogRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexAppCatalogRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexAppCatalogRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexAppCatalogRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexAppCatalogRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexAppCatalogRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexAppCatalogRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexAppCatalogRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexAppCatalogRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexAppCatalogRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexAppCatalogRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexAppCatalogRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexAppCatalogRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexAppCatalogRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexAppCatalogRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexAppCatalogRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetVersion

`func (o *HyperflexAppCatalogRelationship) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HyperflexAppCatalogRelationship) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HyperflexAppCatalogRelationship) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HyperflexAppCatalogRelationship) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatureLimitExternal

`func (o *HyperflexAppCatalogRelationship) GetFeatureLimitExternal() HyperflexFeatureLimitExternalRelationship`

GetFeatureLimitExternal returns the FeatureLimitExternal field if non-nil, zero value otherwise.

### GetFeatureLimitExternalOk

`func (o *HyperflexAppCatalogRelationship) GetFeatureLimitExternalOk() (*HyperflexFeatureLimitExternalRelationship, bool)`

GetFeatureLimitExternalOk returns a tuple with the FeatureLimitExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitExternal

`func (o *HyperflexAppCatalogRelationship) SetFeatureLimitExternal(v HyperflexFeatureLimitExternalRelationship)`

SetFeatureLimitExternal sets FeatureLimitExternal field to given value.

### HasFeatureLimitExternal

`func (o *HyperflexAppCatalogRelationship) HasFeatureLimitExternal() bool`

HasFeatureLimitExternal returns a boolean if a field has been set.

### GetFeatureLimitInternal

`func (o *HyperflexAppCatalogRelationship) GetFeatureLimitInternal() HyperflexFeatureLimitInternalRelationship`

GetFeatureLimitInternal returns the FeatureLimitInternal field if non-nil, zero value otherwise.

### GetFeatureLimitInternalOk

`func (o *HyperflexAppCatalogRelationship) GetFeatureLimitInternalOk() (*HyperflexFeatureLimitInternalRelationship, bool)`

GetFeatureLimitInternalOk returns a tuple with the FeatureLimitInternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureLimitInternal

`func (o *HyperflexAppCatalogRelationship) SetFeatureLimitInternal(v HyperflexFeatureLimitInternalRelationship)`

SetFeatureLimitInternal sets FeatureLimitInternal field to given value.

### HasFeatureLimitInternal

`func (o *HyperflexAppCatalogRelationship) HasFeatureLimitInternal() bool`

HasFeatureLimitInternal returns a boolean if a field has been set.

### GetHxdpVersions

`func (o *HyperflexAppCatalogRelationship) GetHxdpVersions() []HyperflexHxdpVersionRelationship`

GetHxdpVersions returns the HxdpVersions field if non-nil, zero value otherwise.

### GetHxdpVersionsOk

`func (o *HyperflexAppCatalogRelationship) GetHxdpVersionsOk() (*[]HyperflexHxdpVersionRelationship, bool)`

GetHxdpVersionsOk returns a tuple with the HxdpVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpVersions

`func (o *HyperflexAppCatalogRelationship) SetHxdpVersions(v []HyperflexHxdpVersionRelationship)`

SetHxdpVersions sets HxdpVersions field to given value.

### HasHxdpVersions

`func (o *HyperflexAppCatalogRelationship) HasHxdpVersions() bool`

HasHxdpVersions returns a boolean if a field has been set.

### SetHxdpVersionsNil

`func (o *HyperflexAppCatalogRelationship) SetHxdpVersionsNil(b bool)`

 SetHxdpVersionsNil sets the value for HxdpVersions to be an explicit nil

### UnsetHxdpVersions
`func (o *HyperflexAppCatalogRelationship) UnsetHxdpVersions()`

UnsetHxdpVersions ensures that no value is present for HxdpVersions, not even an explicit nil
### GetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogRelationship) GetHyperflexCapabilityInfos() []HyperflexCapabilityInfoRelationship`

GetHyperflexCapabilityInfos returns the HyperflexCapabilityInfos field if non-nil, zero value otherwise.

### GetHyperflexCapabilityInfosOk

`func (o *HyperflexAppCatalogRelationship) GetHyperflexCapabilityInfosOk() (*[]HyperflexCapabilityInfoRelationship, bool)`

GetHyperflexCapabilityInfosOk returns a tuple with the HyperflexCapabilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogRelationship) SetHyperflexCapabilityInfos(v []HyperflexCapabilityInfoRelationship)`

SetHyperflexCapabilityInfos sets HyperflexCapabilityInfos field to given value.

### HasHyperflexCapabilityInfos

`func (o *HyperflexAppCatalogRelationship) HasHyperflexCapabilityInfos() bool`

HasHyperflexCapabilityInfos returns a boolean if a field has been set.

### SetHyperflexCapabilityInfosNil

`func (o *HyperflexAppCatalogRelationship) SetHyperflexCapabilityInfosNil(b bool)`

 SetHyperflexCapabilityInfosNil sets the value for HyperflexCapabilityInfos to be an explicit nil

### UnsetHyperflexCapabilityInfos
`func (o *HyperflexAppCatalogRelationship) UnsetHyperflexCapabilityInfos()`

UnsetHyperflexCapabilityInfos ensures that no value is present for HyperflexCapabilityInfos, not even an explicit nil
### GetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogRelationship) GetHyperflexSoftwareCompatibilityInfos() []HclHyperflexSoftwareCompatibilityInfoRelationship`

GetHyperflexSoftwareCompatibilityInfos returns the HyperflexSoftwareCompatibilityInfos field if non-nil, zero value otherwise.

### GetHyperflexSoftwareCompatibilityInfosOk

`func (o *HyperflexAppCatalogRelationship) GetHyperflexSoftwareCompatibilityInfosOk() (*[]HclHyperflexSoftwareCompatibilityInfoRelationship, bool)`

GetHyperflexSoftwareCompatibilityInfosOk returns a tuple with the HyperflexSoftwareCompatibilityInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogRelationship) SetHyperflexSoftwareCompatibilityInfos(v []HclHyperflexSoftwareCompatibilityInfoRelationship)`

SetHyperflexSoftwareCompatibilityInfos sets HyperflexSoftwareCompatibilityInfos field to given value.

### HasHyperflexSoftwareCompatibilityInfos

`func (o *HyperflexAppCatalogRelationship) HasHyperflexSoftwareCompatibilityInfos() bool`

HasHyperflexSoftwareCompatibilityInfos returns a boolean if a field has been set.

### SetHyperflexSoftwareCompatibilityInfosNil

`func (o *HyperflexAppCatalogRelationship) SetHyperflexSoftwareCompatibilityInfosNil(b bool)`

 SetHyperflexSoftwareCompatibilityInfosNil sets the value for HyperflexSoftwareCompatibilityInfos to be an explicit nil

### UnsetHyperflexSoftwareCompatibilityInfos
`func (o *HyperflexAppCatalogRelationship) UnsetHyperflexSoftwareCompatibilityInfos()`

UnsetHyperflexSoftwareCompatibilityInfos ensures that no value is present for HyperflexSoftwareCompatibilityInfos, not even an explicit nil
### GetServerFirmwareVersion

`func (o *HyperflexAppCatalogRelationship) GetServerFirmwareVersion() HyperflexServerFirmwareVersionRelationship`

GetServerFirmwareVersion returns the ServerFirmwareVersion field if non-nil, zero value otherwise.

### GetServerFirmwareVersionOk

`func (o *HyperflexAppCatalogRelationship) GetServerFirmwareVersionOk() (*HyperflexServerFirmwareVersionRelationship, bool)`

GetServerFirmwareVersionOk returns a tuple with the ServerFirmwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerFirmwareVersion

`func (o *HyperflexAppCatalogRelationship) SetServerFirmwareVersion(v HyperflexServerFirmwareVersionRelationship)`

SetServerFirmwareVersion sets ServerFirmwareVersion field to given value.

### HasServerFirmwareVersion

`func (o *HyperflexAppCatalogRelationship) HasServerFirmwareVersion() bool`

HasServerFirmwareVersion returns a boolean if a field has been set.

### GetServerModel

`func (o *HyperflexAppCatalogRelationship) GetServerModel() HyperflexServerModelRelationship`

GetServerModel returns the ServerModel field if non-nil, zero value otherwise.

### GetServerModelOk

`func (o *HyperflexAppCatalogRelationship) GetServerModelOk() (*HyperflexServerModelRelationship, bool)`

GetServerModelOk returns a tuple with the ServerModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerModel

`func (o *HyperflexAppCatalogRelationship) SetServerModel(v HyperflexServerModelRelationship)`

SetServerModel sets ServerModel field to given value.

### HasServerModel

`func (o *HyperflexAppCatalogRelationship) HasServerModel() bool`

HasServerModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


