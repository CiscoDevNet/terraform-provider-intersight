# OrganizationOrganizationRelationship

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
**Description** | Pointer to **string** | The informative description about the usage of this organization. | [optional] 
**Name** | Pointer to **string** | The name of the organization. There can be multiple organizations under an account. | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**ResourceGroups** | Pointer to [**[]ResourceGroupRelationship**](resource.Group.Relationship.md) | An array of relationships to resourceGroup resources. | [optional] 

## Methods

### NewOrganizationOrganizationRelationship

`func NewOrganizationOrganizationRelationship(classId string, objectType string, ) *OrganizationOrganizationRelationship`

NewOrganizationOrganizationRelationship instantiates a new OrganizationOrganizationRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationOrganizationRelationshipWithDefaults

`func NewOrganizationOrganizationRelationshipWithDefaults() *OrganizationOrganizationRelationship`

NewOrganizationOrganizationRelationshipWithDefaults instantiates a new OrganizationOrganizationRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *OrganizationOrganizationRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *OrganizationOrganizationRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *OrganizationOrganizationRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *OrganizationOrganizationRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *OrganizationOrganizationRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *OrganizationOrganizationRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *OrganizationOrganizationRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *OrganizationOrganizationRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *OrganizationOrganizationRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *OrganizationOrganizationRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *OrganizationOrganizationRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *OrganizationOrganizationRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *OrganizationOrganizationRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *OrganizationOrganizationRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *OrganizationOrganizationRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *OrganizationOrganizationRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *OrganizationOrganizationRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *OrganizationOrganizationRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *OrganizationOrganizationRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *OrganizationOrganizationRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *OrganizationOrganizationRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *OrganizationOrganizationRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *OrganizationOrganizationRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *OrganizationOrganizationRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *OrganizationOrganizationRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *OrganizationOrganizationRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *OrganizationOrganizationRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *OrganizationOrganizationRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *OrganizationOrganizationRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *OrganizationOrganizationRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *OrganizationOrganizationRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *OrganizationOrganizationRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *OrganizationOrganizationRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *OrganizationOrganizationRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *OrganizationOrganizationRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *OrganizationOrganizationRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *OrganizationOrganizationRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *OrganizationOrganizationRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *OrganizationOrganizationRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *OrganizationOrganizationRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *OrganizationOrganizationRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *OrganizationOrganizationRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *OrganizationOrganizationRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *OrganizationOrganizationRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *OrganizationOrganizationRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *OrganizationOrganizationRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *OrganizationOrganizationRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *OrganizationOrganizationRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *OrganizationOrganizationRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *OrganizationOrganizationRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *OrganizationOrganizationRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *OrganizationOrganizationRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *OrganizationOrganizationRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *OrganizationOrganizationRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *OrganizationOrganizationRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *OrganizationOrganizationRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *OrganizationOrganizationRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *OrganizationOrganizationRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *OrganizationOrganizationRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *OrganizationOrganizationRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *OrganizationOrganizationRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *OrganizationOrganizationRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *OrganizationOrganizationRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *OrganizationOrganizationRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *OrganizationOrganizationRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrganizationOrganizationRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrganizationOrganizationRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrganizationOrganizationRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *OrganizationOrganizationRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationOrganizationRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationOrganizationRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrganizationOrganizationRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccount

`func (o *OrganizationOrganizationRelationship) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OrganizationOrganizationRelationship) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OrganizationOrganizationRelationship) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OrganizationOrganizationRelationship) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetResourceGroups

`func (o *OrganizationOrganizationRelationship) GetResourceGroups() []ResourceGroupRelationship`

GetResourceGroups returns the ResourceGroups field if non-nil, zero value otherwise.

### GetResourceGroupsOk

`func (o *OrganizationOrganizationRelationship) GetResourceGroupsOk() (*[]ResourceGroupRelationship, bool)`

GetResourceGroupsOk returns a tuple with the ResourceGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroups

`func (o *OrganizationOrganizationRelationship) SetResourceGroups(v []ResourceGroupRelationship)`

SetResourceGroups sets ResourceGroups field to given value.

### HasResourceGroups

`func (o *OrganizationOrganizationRelationship) HasResourceGroups() bool`

HasResourceGroups returns a boolean if a field has been set.

### SetResourceGroupsNil

`func (o *OrganizationOrganizationRelationship) SetResourceGroupsNil(b bool)`

 SetResourceGroupsNil sets the value for ResourceGroups to be an explicit nil

### UnsetResourceGroups
`func (o *OrganizationOrganizationRelationship) UnsetResourceGroups()`

UnsetResourceGroups ensures that no value is present for ResourceGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


