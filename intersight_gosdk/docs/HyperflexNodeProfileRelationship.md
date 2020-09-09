# HyperflexNodeProfileRelationship

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
**Description** | Pointer to **string** | Description of the profile. | [optional] 
**Name** | Pointer to **string** | Name of the concrete profile. | [optional] 
**Type** | Pointer to **string** | Defines the type of the profile. Accepted value is instance. * &#x60;instance&#x60; - The profile defines the configuration for a specific instance of a target. | [optional] [default to "instance"]
**SrcTemplate** | Pointer to [**PolicyAbstractProfileRelationship**](policy.AbstractProfile.Relationship.md) |  | [optional] 
**HxdpDataIp** | Pointer to **string** | IP address for storage data network (Controller VM interface). | [optional] 
**HxdpMgmtIp** | Pointer to **string** | IP address for HyperFlex management network. | [optional] 
**HypervisorDataIp** | Pointer to **string** | IP address for storage data network (Hypervisor interface). | [optional] 
**HypervisorMgmtIp** | Pointer to **string** | IP address for Hypervisor management network. | [optional] 
**AssignedServer** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**ClusterProfile** | Pointer to [**HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexNodeProfileRelationship

`func NewHyperflexNodeProfileRelationship(classId string, objectType string, ) *HyperflexNodeProfileRelationship`

NewHyperflexNodeProfileRelationship instantiates a new HyperflexNodeProfileRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexNodeProfileRelationshipWithDefaults

`func NewHyperflexNodeProfileRelationshipWithDefaults() *HyperflexNodeProfileRelationship`

NewHyperflexNodeProfileRelationshipWithDefaults instantiates a new HyperflexNodeProfileRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexNodeProfileRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexNodeProfileRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexNodeProfileRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexNodeProfileRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexNodeProfileRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexNodeProfileRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexNodeProfileRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexNodeProfileRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexNodeProfileRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexNodeProfileRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexNodeProfileRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexNodeProfileRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexNodeProfileRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexNodeProfileRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexNodeProfileRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexNodeProfileRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexNodeProfileRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexNodeProfileRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexNodeProfileRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexNodeProfileRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexNodeProfileRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexNodeProfileRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexNodeProfileRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexNodeProfileRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexNodeProfileRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexNodeProfileRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexNodeProfileRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexNodeProfileRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexNodeProfileRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexNodeProfileRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexNodeProfileRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexNodeProfileRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexNodeProfileRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexNodeProfileRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexNodeProfileRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexNodeProfileRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexNodeProfileRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexNodeProfileRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexNodeProfileRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexNodeProfileRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexNodeProfileRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexNodeProfileRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexNodeProfileRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexNodeProfileRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexNodeProfileRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexNodeProfileRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexNodeProfileRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexNodeProfileRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexNodeProfileRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexNodeProfileRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexNodeProfileRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexNodeProfileRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexNodeProfileRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexNodeProfileRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexNodeProfileRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexNodeProfileRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexNodeProfileRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexNodeProfileRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexNodeProfileRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexNodeProfileRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexNodeProfileRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexNodeProfileRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexNodeProfileRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexNodeProfileRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *HyperflexNodeProfileRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexNodeProfileRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexNodeProfileRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexNodeProfileRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HyperflexNodeProfileRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexNodeProfileRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexNodeProfileRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexNodeProfileRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *HyperflexNodeProfileRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexNodeProfileRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexNodeProfileRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexNodeProfileRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSrcTemplate

`func (o *HyperflexNodeProfileRelationship) GetSrcTemplate() PolicyAbstractProfileRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *HyperflexNodeProfileRelationship) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *HyperflexNodeProfileRelationship) SetSrcTemplate(v PolicyAbstractProfileRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *HyperflexNodeProfileRelationship) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.

### GetHxdpDataIp

`func (o *HyperflexNodeProfileRelationship) GetHxdpDataIp() string`

GetHxdpDataIp returns the HxdpDataIp field if non-nil, zero value otherwise.

### GetHxdpDataIpOk

`func (o *HyperflexNodeProfileRelationship) GetHxdpDataIpOk() (*string, bool)`

GetHxdpDataIpOk returns a tuple with the HxdpDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpDataIp

`func (o *HyperflexNodeProfileRelationship) SetHxdpDataIp(v string)`

SetHxdpDataIp sets HxdpDataIp field to given value.

### HasHxdpDataIp

`func (o *HyperflexNodeProfileRelationship) HasHxdpDataIp() bool`

HasHxdpDataIp returns a boolean if a field has been set.

### GetHxdpMgmtIp

`func (o *HyperflexNodeProfileRelationship) GetHxdpMgmtIp() string`

GetHxdpMgmtIp returns the HxdpMgmtIp field if non-nil, zero value otherwise.

### GetHxdpMgmtIpOk

`func (o *HyperflexNodeProfileRelationship) GetHxdpMgmtIpOk() (*string, bool)`

GetHxdpMgmtIpOk returns a tuple with the HxdpMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpMgmtIp

`func (o *HyperflexNodeProfileRelationship) SetHxdpMgmtIp(v string)`

SetHxdpMgmtIp sets HxdpMgmtIp field to given value.

### HasHxdpMgmtIp

`func (o *HyperflexNodeProfileRelationship) HasHxdpMgmtIp() bool`

HasHxdpMgmtIp returns a boolean if a field has been set.

### GetHypervisorDataIp

`func (o *HyperflexNodeProfileRelationship) GetHypervisorDataIp() string`

GetHypervisorDataIp returns the HypervisorDataIp field if non-nil, zero value otherwise.

### GetHypervisorDataIpOk

`func (o *HyperflexNodeProfileRelationship) GetHypervisorDataIpOk() (*string, bool)`

GetHypervisorDataIpOk returns a tuple with the HypervisorDataIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorDataIp

`func (o *HyperflexNodeProfileRelationship) SetHypervisorDataIp(v string)`

SetHypervisorDataIp sets HypervisorDataIp field to given value.

### HasHypervisorDataIp

`func (o *HyperflexNodeProfileRelationship) HasHypervisorDataIp() bool`

HasHypervisorDataIp returns a boolean if a field has been set.

### GetHypervisorMgmtIp

`func (o *HyperflexNodeProfileRelationship) GetHypervisorMgmtIp() string`

GetHypervisorMgmtIp returns the HypervisorMgmtIp field if non-nil, zero value otherwise.

### GetHypervisorMgmtIpOk

`func (o *HyperflexNodeProfileRelationship) GetHypervisorMgmtIpOk() (*string, bool)`

GetHypervisorMgmtIpOk returns a tuple with the HypervisorMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorMgmtIp

`func (o *HyperflexNodeProfileRelationship) SetHypervisorMgmtIp(v string)`

SetHypervisorMgmtIp sets HypervisorMgmtIp field to given value.

### HasHypervisorMgmtIp

`func (o *HyperflexNodeProfileRelationship) HasHypervisorMgmtIp() bool`

HasHypervisorMgmtIp returns a boolean if a field has been set.

### GetAssignedServer

`func (o *HyperflexNodeProfileRelationship) GetAssignedServer() ComputeRackUnitRelationship`

GetAssignedServer returns the AssignedServer field if non-nil, zero value otherwise.

### GetAssignedServerOk

`func (o *HyperflexNodeProfileRelationship) GetAssignedServerOk() (*ComputeRackUnitRelationship, bool)`

GetAssignedServerOk returns a tuple with the AssignedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedServer

`func (o *HyperflexNodeProfileRelationship) SetAssignedServer(v ComputeRackUnitRelationship)`

SetAssignedServer sets AssignedServer field to given value.

### HasAssignedServer

`func (o *HyperflexNodeProfileRelationship) HasAssignedServer() bool`

HasAssignedServer returns a boolean if a field has been set.

### GetClusterProfile

`func (o *HyperflexNodeProfileRelationship) GetClusterProfile() HyperflexClusterProfileRelationship`

GetClusterProfile returns the ClusterProfile field if non-nil, zero value otherwise.

### GetClusterProfileOk

`func (o *HyperflexNodeProfileRelationship) GetClusterProfileOk() (*HyperflexClusterProfileRelationship, bool)`

GetClusterProfileOk returns a tuple with the ClusterProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfile

`func (o *HyperflexNodeProfileRelationship) SetClusterProfile(v HyperflexClusterProfileRelationship)`

SetClusterProfile sets ClusterProfile field to given value.

### HasClusterProfile

`func (o *HyperflexNodeProfileRelationship) HasClusterProfile() bool`

HasClusterProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


