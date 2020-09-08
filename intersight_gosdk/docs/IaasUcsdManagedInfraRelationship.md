# IaasUcsdManagedInfraRelationship

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
**AdvancedCatalogCount** | Pointer to **int64** | Total advanced catalogs in UCSD. | [optional] [readonly] 
**BmCatalogCount** | Pointer to **int64** | Total bare metal catalogs in UCSD. | [optional] [readonly] 
**ContainerCatalogCount** | Pointer to **int64** | Total service container catalogs in UCSD. | [optional] [readonly] 
**EsxiHostCount** | Pointer to **int64** | Total ESXi hosts in UCSD. | [optional] [readonly] 
**ExternalGroupCount** | Pointer to **int64** | Total external (Ldap) groups in UCSD. | [optional] [readonly] 
**HypervHostCount** | Pointer to **int64** | Total HyperV hosts in UCSD. | [optional] [readonly] 
**LocalGroupCount** | Pointer to **int64** | Total local groups in UCSD. | [optional] [readonly] 
**StandardCatalogCount** | Pointer to **int64** | Total standard catalogs in UCSD. | [optional] [readonly] 
**UserCount** | Pointer to **int64** | Total user accounts in UCSD. | [optional] [readonly] 
**VdcCount** | Pointer to **int64** | Total virtual datacenters in UCSD. | [optional] [readonly] 
**VmCount** | Pointer to **int64** | Total Virtual machines in UCSD. | [optional] [readonly] 
**Guid** | Pointer to [**IaasUcsdInfoRelationship**](iaas.UcsdInfo.Relationship.md) |  | [optional] 

## Methods

### NewIaasUcsdManagedInfraRelationship

`func NewIaasUcsdManagedInfraRelationship(classId string, objectType string, ) *IaasUcsdManagedInfraRelationship`

NewIaasUcsdManagedInfraRelationship instantiates a new IaasUcsdManagedInfraRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIaasUcsdManagedInfraRelationshipWithDefaults

`func NewIaasUcsdManagedInfraRelationshipWithDefaults() *IaasUcsdManagedInfraRelationship`

NewIaasUcsdManagedInfraRelationshipWithDefaults instantiates a new IaasUcsdManagedInfraRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *IaasUcsdManagedInfraRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *IaasUcsdManagedInfraRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *IaasUcsdManagedInfraRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *IaasUcsdManagedInfraRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *IaasUcsdManagedInfraRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *IaasUcsdManagedInfraRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *IaasUcsdManagedInfraRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *IaasUcsdManagedInfraRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *IaasUcsdManagedInfraRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *IaasUcsdManagedInfraRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *IaasUcsdManagedInfraRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *IaasUcsdManagedInfraRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *IaasUcsdManagedInfraRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *IaasUcsdManagedInfraRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *IaasUcsdManagedInfraRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *IaasUcsdManagedInfraRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *IaasUcsdManagedInfraRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *IaasUcsdManagedInfraRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *IaasUcsdManagedInfraRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *IaasUcsdManagedInfraRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *IaasUcsdManagedInfraRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *IaasUcsdManagedInfraRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *IaasUcsdManagedInfraRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *IaasUcsdManagedInfraRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *IaasUcsdManagedInfraRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *IaasUcsdManagedInfraRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *IaasUcsdManagedInfraRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *IaasUcsdManagedInfraRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *IaasUcsdManagedInfraRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *IaasUcsdManagedInfraRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *IaasUcsdManagedInfraRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *IaasUcsdManagedInfraRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *IaasUcsdManagedInfraRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *IaasUcsdManagedInfraRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *IaasUcsdManagedInfraRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IaasUcsdManagedInfraRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IaasUcsdManagedInfraRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *IaasUcsdManagedInfraRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *IaasUcsdManagedInfraRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *IaasUcsdManagedInfraRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *IaasUcsdManagedInfraRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *IaasUcsdManagedInfraRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *IaasUcsdManagedInfraRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *IaasUcsdManagedInfraRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *IaasUcsdManagedInfraRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *IaasUcsdManagedInfraRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *IaasUcsdManagedInfraRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *IaasUcsdManagedInfraRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *IaasUcsdManagedInfraRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IaasUcsdManagedInfraRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IaasUcsdManagedInfraRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IaasUcsdManagedInfraRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *IaasUcsdManagedInfraRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *IaasUcsdManagedInfraRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *IaasUcsdManagedInfraRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *IaasUcsdManagedInfraRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *IaasUcsdManagedInfraRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *IaasUcsdManagedInfraRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *IaasUcsdManagedInfraRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *IaasUcsdManagedInfraRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *IaasUcsdManagedInfraRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *IaasUcsdManagedInfraRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *IaasUcsdManagedInfraRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *IaasUcsdManagedInfraRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) GetAdvancedCatalogCount() int64`

GetAdvancedCatalogCount returns the AdvancedCatalogCount field if non-nil, zero value otherwise.

### GetAdvancedCatalogCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetAdvancedCatalogCountOk() (*int64, bool)`

GetAdvancedCatalogCountOk returns a tuple with the AdvancedCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) SetAdvancedCatalogCount(v int64)`

SetAdvancedCatalogCount sets AdvancedCatalogCount field to given value.

### HasAdvancedCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) HasAdvancedCatalogCount() bool`

HasAdvancedCatalogCount returns a boolean if a field has been set.

### GetBmCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) GetBmCatalogCount() int64`

GetBmCatalogCount returns the BmCatalogCount field if non-nil, zero value otherwise.

### GetBmCatalogCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetBmCatalogCountOk() (*int64, bool)`

GetBmCatalogCountOk returns a tuple with the BmCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBmCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) SetBmCatalogCount(v int64)`

SetBmCatalogCount sets BmCatalogCount field to given value.

### HasBmCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) HasBmCatalogCount() bool`

HasBmCatalogCount returns a boolean if a field has been set.

### GetContainerCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) GetContainerCatalogCount() int64`

GetContainerCatalogCount returns the ContainerCatalogCount field if non-nil, zero value otherwise.

### GetContainerCatalogCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetContainerCatalogCountOk() (*int64, bool)`

GetContainerCatalogCountOk returns a tuple with the ContainerCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) SetContainerCatalogCount(v int64)`

SetContainerCatalogCount sets ContainerCatalogCount field to given value.

### HasContainerCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) HasContainerCatalogCount() bool`

HasContainerCatalogCount returns a boolean if a field has been set.

### GetEsxiHostCount

`func (o *IaasUcsdManagedInfraRelationship) GetEsxiHostCount() int64`

GetEsxiHostCount returns the EsxiHostCount field if non-nil, zero value otherwise.

### GetEsxiHostCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetEsxiHostCountOk() (*int64, bool)`

GetEsxiHostCountOk returns a tuple with the EsxiHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsxiHostCount

`func (o *IaasUcsdManagedInfraRelationship) SetEsxiHostCount(v int64)`

SetEsxiHostCount sets EsxiHostCount field to given value.

### HasEsxiHostCount

`func (o *IaasUcsdManagedInfraRelationship) HasEsxiHostCount() bool`

HasEsxiHostCount returns a boolean if a field has been set.

### GetExternalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) GetExternalGroupCount() int64`

GetExternalGroupCount returns the ExternalGroupCount field if non-nil, zero value otherwise.

### GetExternalGroupCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetExternalGroupCountOk() (*int64, bool)`

GetExternalGroupCountOk returns a tuple with the ExternalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) SetExternalGroupCount(v int64)`

SetExternalGroupCount sets ExternalGroupCount field to given value.

### HasExternalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) HasExternalGroupCount() bool`

HasExternalGroupCount returns a boolean if a field has been set.

### GetHypervHostCount

`func (o *IaasUcsdManagedInfraRelationship) GetHypervHostCount() int64`

GetHypervHostCount returns the HypervHostCount field if non-nil, zero value otherwise.

### GetHypervHostCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetHypervHostCountOk() (*int64, bool)`

GetHypervHostCountOk returns a tuple with the HypervHostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervHostCount

`func (o *IaasUcsdManagedInfraRelationship) SetHypervHostCount(v int64)`

SetHypervHostCount sets HypervHostCount field to given value.

### HasHypervHostCount

`func (o *IaasUcsdManagedInfraRelationship) HasHypervHostCount() bool`

HasHypervHostCount returns a boolean if a field has been set.

### GetLocalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) GetLocalGroupCount() int64`

GetLocalGroupCount returns the LocalGroupCount field if non-nil, zero value otherwise.

### GetLocalGroupCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetLocalGroupCountOk() (*int64, bool)`

GetLocalGroupCountOk returns a tuple with the LocalGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) SetLocalGroupCount(v int64)`

SetLocalGroupCount sets LocalGroupCount field to given value.

### HasLocalGroupCount

`func (o *IaasUcsdManagedInfraRelationship) HasLocalGroupCount() bool`

HasLocalGroupCount returns a boolean if a field has been set.

### GetStandardCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) GetStandardCatalogCount() int64`

GetStandardCatalogCount returns the StandardCatalogCount field if non-nil, zero value otherwise.

### GetStandardCatalogCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetStandardCatalogCountOk() (*int64, bool)`

GetStandardCatalogCountOk returns a tuple with the StandardCatalogCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) SetStandardCatalogCount(v int64)`

SetStandardCatalogCount sets StandardCatalogCount field to given value.

### HasStandardCatalogCount

`func (o *IaasUcsdManagedInfraRelationship) HasStandardCatalogCount() bool`

HasStandardCatalogCount returns a boolean if a field has been set.

### GetUserCount

`func (o *IaasUcsdManagedInfraRelationship) GetUserCount() int64`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetUserCountOk() (*int64, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *IaasUcsdManagedInfraRelationship) SetUserCount(v int64)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *IaasUcsdManagedInfraRelationship) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetVdcCount

`func (o *IaasUcsdManagedInfraRelationship) GetVdcCount() int64`

GetVdcCount returns the VdcCount field if non-nil, zero value otherwise.

### GetVdcCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetVdcCountOk() (*int64, bool)`

GetVdcCountOk returns a tuple with the VdcCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcCount

`func (o *IaasUcsdManagedInfraRelationship) SetVdcCount(v int64)`

SetVdcCount sets VdcCount field to given value.

### HasVdcCount

`func (o *IaasUcsdManagedInfraRelationship) HasVdcCount() bool`

HasVdcCount returns a boolean if a field has been set.

### GetVmCount

`func (o *IaasUcsdManagedInfraRelationship) GetVmCount() int64`

GetVmCount returns the VmCount field if non-nil, zero value otherwise.

### GetVmCountOk

`func (o *IaasUcsdManagedInfraRelationship) GetVmCountOk() (*int64, bool)`

GetVmCountOk returns a tuple with the VmCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCount

`func (o *IaasUcsdManagedInfraRelationship) SetVmCount(v int64)`

SetVmCount sets VmCount field to given value.

### HasVmCount

`func (o *IaasUcsdManagedInfraRelationship) HasVmCount() bool`

HasVmCount returns a boolean if a field has been set.

### GetGuid

`func (o *IaasUcsdManagedInfraRelationship) GetGuid() IaasUcsdInfoRelationship`

GetGuid returns the Guid field if non-nil, zero value otherwise.

### GetGuidOk

`func (o *IaasUcsdManagedInfraRelationship) GetGuidOk() (*IaasUcsdInfoRelationship, bool)`

GetGuidOk returns a tuple with the Guid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuid

`func (o *IaasUcsdManagedInfraRelationship) SetGuid(v IaasUcsdInfoRelationship)`

SetGuid sets Guid field to given value.

### HasGuid

`func (o *IaasUcsdManagedInfraRelationship) HasGuid() bool`

HasGuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


