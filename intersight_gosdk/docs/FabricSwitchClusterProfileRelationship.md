# FabricSwitchClusterProfileRelationship

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
**ConfigContext** | Pointer to [**PolicyConfigContext**](policy.ConfigContext.md) |  | [optional] 
**SwitchProfilesCount** | Pointer to **int64** | Number of switch profiles that are part of this cluster profile. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**SwitchProfiles** | Pointer to [**[]FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricSwitchClusterProfileRelationship

`func NewFabricSwitchClusterProfileRelationship(classId string, objectType string, ) *FabricSwitchClusterProfileRelationship`

NewFabricSwitchClusterProfileRelationship instantiates a new FabricSwitchClusterProfileRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricSwitchClusterProfileRelationshipWithDefaults

`func NewFabricSwitchClusterProfileRelationshipWithDefaults() *FabricSwitchClusterProfileRelationship`

NewFabricSwitchClusterProfileRelationshipWithDefaults instantiates a new FabricSwitchClusterProfileRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FabricSwitchClusterProfileRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FabricSwitchClusterProfileRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FabricSwitchClusterProfileRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FabricSwitchClusterProfileRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FabricSwitchClusterProfileRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricSwitchClusterProfileRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricSwitchClusterProfileRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FabricSwitchClusterProfileRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FabricSwitchClusterProfileRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FabricSwitchClusterProfileRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FabricSwitchClusterProfileRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FabricSwitchClusterProfileRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FabricSwitchClusterProfileRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FabricSwitchClusterProfileRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FabricSwitchClusterProfileRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FabricSwitchClusterProfileRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FabricSwitchClusterProfileRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FabricSwitchClusterProfileRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FabricSwitchClusterProfileRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FabricSwitchClusterProfileRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FabricSwitchClusterProfileRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FabricSwitchClusterProfileRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FabricSwitchClusterProfileRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FabricSwitchClusterProfileRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricSwitchClusterProfileRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricSwitchClusterProfileRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FabricSwitchClusterProfileRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FabricSwitchClusterProfileRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FabricSwitchClusterProfileRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FabricSwitchClusterProfileRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FabricSwitchClusterProfileRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FabricSwitchClusterProfileRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FabricSwitchClusterProfileRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FabricSwitchClusterProfileRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FabricSwitchClusterProfileRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FabricSwitchClusterProfileRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FabricSwitchClusterProfileRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FabricSwitchClusterProfileRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FabricSwitchClusterProfileRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FabricSwitchClusterProfileRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FabricSwitchClusterProfileRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FabricSwitchClusterProfileRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FabricSwitchClusterProfileRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FabricSwitchClusterProfileRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FabricSwitchClusterProfileRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FabricSwitchClusterProfileRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FabricSwitchClusterProfileRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FabricSwitchClusterProfileRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FabricSwitchClusterProfileRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FabricSwitchClusterProfileRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FabricSwitchClusterProfileRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FabricSwitchClusterProfileRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FabricSwitchClusterProfileRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FabricSwitchClusterProfileRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FabricSwitchClusterProfileRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FabricSwitchClusterProfileRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FabricSwitchClusterProfileRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FabricSwitchClusterProfileRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FabricSwitchClusterProfileRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FabricSwitchClusterProfileRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FabricSwitchClusterProfileRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FabricSwitchClusterProfileRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FabricSwitchClusterProfileRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FabricSwitchClusterProfileRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *FabricSwitchClusterProfileRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FabricSwitchClusterProfileRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FabricSwitchClusterProfileRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FabricSwitchClusterProfileRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FabricSwitchClusterProfileRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FabricSwitchClusterProfileRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FabricSwitchClusterProfileRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FabricSwitchClusterProfileRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *FabricSwitchClusterProfileRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FabricSwitchClusterProfileRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FabricSwitchClusterProfileRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FabricSwitchClusterProfileRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSrcTemplate

`func (o *FabricSwitchClusterProfileRelationship) GetSrcTemplate() PolicyAbstractProfileRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *FabricSwitchClusterProfileRelationship) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *FabricSwitchClusterProfileRelationship) SetSrcTemplate(v PolicyAbstractProfileRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *FabricSwitchClusterProfileRelationship) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.

### GetConfigContext

`func (o *FabricSwitchClusterProfileRelationship) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *FabricSwitchClusterProfileRelationship) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *FabricSwitchClusterProfileRelationship) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *FabricSwitchClusterProfileRelationship) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetSwitchProfilesCount

`func (o *FabricSwitchClusterProfileRelationship) GetSwitchProfilesCount() int64`

GetSwitchProfilesCount returns the SwitchProfilesCount field if non-nil, zero value otherwise.

### GetSwitchProfilesCountOk

`func (o *FabricSwitchClusterProfileRelationship) GetSwitchProfilesCountOk() (*int64, bool)`

GetSwitchProfilesCountOk returns a tuple with the SwitchProfilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfilesCount

`func (o *FabricSwitchClusterProfileRelationship) SetSwitchProfilesCount(v int64)`

SetSwitchProfilesCount sets SwitchProfilesCount field to given value.

### HasSwitchProfilesCount

`func (o *FabricSwitchClusterProfileRelationship) HasSwitchProfilesCount() bool`

HasSwitchProfilesCount returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricSwitchClusterProfileRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricSwitchClusterProfileRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricSwitchClusterProfileRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricSwitchClusterProfileRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSwitchProfiles

`func (o *FabricSwitchClusterProfileRelationship) GetSwitchProfiles() []FabricSwitchProfileRelationship`

GetSwitchProfiles returns the SwitchProfiles field if non-nil, zero value otherwise.

### GetSwitchProfilesOk

`func (o *FabricSwitchClusterProfileRelationship) GetSwitchProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetSwitchProfilesOk returns a tuple with the SwitchProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchProfiles

`func (o *FabricSwitchClusterProfileRelationship) SetSwitchProfiles(v []FabricSwitchProfileRelationship)`

SetSwitchProfiles sets SwitchProfiles field to given value.

### HasSwitchProfiles

`func (o *FabricSwitchClusterProfileRelationship) HasSwitchProfiles() bool`

HasSwitchProfiles returns a boolean if a field has been set.

### SetSwitchProfilesNil

`func (o *FabricSwitchClusterProfileRelationship) SetSwitchProfilesNil(b bool)`

 SetSwitchProfilesNil sets the value for SwitchProfiles to be an explicit nil

### UnsetSwitchProfiles
`func (o *FabricSwitchClusterProfileRelationship) UnsetSwitchProfiles()`

UnsetSwitchProfiles ensures that no value is present for SwitchProfiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


