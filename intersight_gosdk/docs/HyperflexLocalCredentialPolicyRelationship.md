# HyperflexLocalCredentialPolicyRelationship

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
**Description** | Pointer to **string** | Description of the policy. | [optional] 
**Name** | Pointer to **string** | Name of the concrete policy. | [optional] 
**FactoryHypervisorPassword** | Pointer to **bool** | Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment. | [optional] 
**HxdpRootPwd** | Pointer to **string** | HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&amp;*! special characters. | [optional] 
**HypervisorAdmin** | Pointer to **string** | Hypervisor administrator username must contain only alphanumeric characters. Use the root account for ESXi deployments. | [optional] 
**HypervisorAdminPwd** | Pointer to **string** | The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password \&quot;Cisco123\&quot; for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment. | [optional] 
**IsHxdpRootPwdSet** | Pointer to **bool** | Indicates whether the value of the &#39;hxdpRootPwd&#39; property has been set. | [optional] [readonly] 
**IsHypervisorAdminPwdSet** | Pointer to **bool** | Indicates whether the value of the &#39;hypervisorAdminPwd&#39; property has been set. | [optional] [readonly] 
**ClusterProfiles** | Pointer to [**[]HyperflexClusterProfileRelationship**](hyperflex.ClusterProfile.Relationship.md) | An array of relationships to hyperflexClusterProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexLocalCredentialPolicyRelationship

`func NewHyperflexLocalCredentialPolicyRelationship(classId string, objectType string, ) *HyperflexLocalCredentialPolicyRelationship`

NewHyperflexLocalCredentialPolicyRelationship instantiates a new HyperflexLocalCredentialPolicyRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexLocalCredentialPolicyRelationshipWithDefaults

`func NewHyperflexLocalCredentialPolicyRelationshipWithDefaults() *HyperflexLocalCredentialPolicyRelationship`

NewHyperflexLocalCredentialPolicyRelationshipWithDefaults instantiates a new HyperflexLocalCredentialPolicyRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexLocalCredentialPolicyRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexLocalCredentialPolicyRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexLocalCredentialPolicyRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexLocalCredentialPolicyRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexLocalCredentialPolicyRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexLocalCredentialPolicyRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexLocalCredentialPolicyRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexLocalCredentialPolicyRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexLocalCredentialPolicyRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexLocalCredentialPolicyRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexLocalCredentialPolicyRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexLocalCredentialPolicyRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexLocalCredentialPolicyRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexLocalCredentialPolicyRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexLocalCredentialPolicyRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexLocalCredentialPolicyRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexLocalCredentialPolicyRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexLocalCredentialPolicyRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexLocalCredentialPolicyRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexLocalCredentialPolicyRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexLocalCredentialPolicyRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexLocalCredentialPolicyRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexLocalCredentialPolicyRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexLocalCredentialPolicyRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexLocalCredentialPolicyRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexLocalCredentialPolicyRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexLocalCredentialPolicyRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexLocalCredentialPolicyRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexLocalCredentialPolicyRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexLocalCredentialPolicyRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexLocalCredentialPolicyRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexLocalCredentialPolicyRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexLocalCredentialPolicyRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexLocalCredentialPolicyRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexLocalCredentialPolicyRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexLocalCredentialPolicyRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexLocalCredentialPolicyRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexLocalCredentialPolicyRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexLocalCredentialPolicyRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexLocalCredentialPolicyRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexLocalCredentialPolicyRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexLocalCredentialPolicyRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HyperflexLocalCredentialPolicyRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexLocalCredentialPolicyRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexLocalCredentialPolicyRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicyRelationship) GetFactoryHypervisorPassword() bool`

GetFactoryHypervisorPassword returns the FactoryHypervisorPassword field if non-nil, zero value otherwise.

### GetFactoryHypervisorPasswordOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetFactoryHypervisorPasswordOk() (*bool, bool)`

GetFactoryHypervisorPasswordOk returns a tuple with the FactoryHypervisorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicyRelationship) SetFactoryHypervisorPassword(v bool)`

SetFactoryHypervisorPassword sets FactoryHypervisorPassword field to given value.

### HasFactoryHypervisorPassword

`func (o *HyperflexLocalCredentialPolicyRelationship) HasFactoryHypervisorPassword() bool`

HasFactoryHypervisorPassword returns a boolean if a field has been set.

### GetHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHxdpRootPwd() string`

GetHxdpRootPwd returns the HxdpRootPwd field if non-nil, zero value otherwise.

### GetHxdpRootPwdOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHxdpRootPwdOk() (*string, bool)`

GetHxdpRootPwdOk returns a tuple with the HxdpRootPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) SetHxdpRootPwd(v string)`

SetHxdpRootPwd sets HxdpRootPwd field to given value.

### HasHxdpRootPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) HasHxdpRootPwd() bool`

HasHxdpRootPwd returns a boolean if a field has been set.

### GetHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHypervisorAdmin() string`

GetHypervisorAdmin returns the HypervisorAdmin field if non-nil, zero value otherwise.

### GetHypervisorAdminOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHypervisorAdminOk() (*string, bool)`

GetHypervisorAdminOk returns a tuple with the HypervisorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicyRelationship) SetHypervisorAdmin(v string)`

SetHypervisorAdmin sets HypervisorAdmin field to given value.

### HasHypervisorAdmin

`func (o *HyperflexLocalCredentialPolicyRelationship) HasHypervisorAdmin() bool`

HasHypervisorAdmin returns a boolean if a field has been set.

### GetHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHypervisorAdminPwd() string`

GetHypervisorAdminPwd returns the HypervisorAdminPwd field if non-nil, zero value otherwise.

### GetHypervisorAdminPwdOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetHypervisorAdminPwdOk() (*string, bool)`

GetHypervisorAdminPwdOk returns a tuple with the HypervisorAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) SetHypervisorAdminPwd(v string)`

SetHypervisorAdminPwd sets HypervisorAdminPwd field to given value.

### HasHypervisorAdminPwd

`func (o *HyperflexLocalCredentialPolicyRelationship) HasHypervisorAdminPwd() bool`

HasHypervisorAdminPwd returns a boolean if a field has been set.

### GetIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) GetIsHxdpRootPwdSet() bool`

GetIsHxdpRootPwdSet returns the IsHxdpRootPwdSet field if non-nil, zero value otherwise.

### GetIsHxdpRootPwdSetOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetIsHxdpRootPwdSetOk() (*bool, bool)`

GetIsHxdpRootPwdSetOk returns a tuple with the IsHxdpRootPwdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) SetIsHxdpRootPwdSet(v bool)`

SetIsHxdpRootPwdSet sets IsHxdpRootPwdSet field to given value.

### HasIsHxdpRootPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) HasIsHxdpRootPwdSet() bool`

HasIsHxdpRootPwdSet returns a boolean if a field has been set.

### GetIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) GetIsHypervisorAdminPwdSet() bool`

GetIsHypervisorAdminPwdSet returns the IsHypervisorAdminPwdSet field if non-nil, zero value otherwise.

### GetIsHypervisorAdminPwdSetOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetIsHypervisorAdminPwdSetOk() (*bool, bool)`

GetIsHypervisorAdminPwdSetOk returns a tuple with the IsHypervisorAdminPwdSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) SetIsHypervisorAdminPwdSet(v bool)`

SetIsHypervisorAdminPwdSet sets IsHypervisorAdminPwdSet field to given value.

### HasIsHypervisorAdminPwdSet

`func (o *HyperflexLocalCredentialPolicyRelationship) HasIsHypervisorAdminPwdSet() bool`

HasIsHypervisorAdminPwdSet returns a boolean if a field has been set.

### GetClusterProfiles

`func (o *HyperflexLocalCredentialPolicyRelationship) GetClusterProfiles() []HyperflexClusterProfileRelationship`

GetClusterProfiles returns the ClusterProfiles field if non-nil, zero value otherwise.

### GetClusterProfilesOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetClusterProfilesOk() (*[]HyperflexClusterProfileRelationship, bool)`

GetClusterProfilesOk returns a tuple with the ClusterProfiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterProfiles

`func (o *HyperflexLocalCredentialPolicyRelationship) SetClusterProfiles(v []HyperflexClusterProfileRelationship)`

SetClusterProfiles sets ClusterProfiles field to given value.

### HasClusterProfiles

`func (o *HyperflexLocalCredentialPolicyRelationship) HasClusterProfiles() bool`

HasClusterProfiles returns a boolean if a field has been set.

### SetClusterProfilesNil

`func (o *HyperflexLocalCredentialPolicyRelationship) SetClusterProfilesNil(b bool)`

 SetClusterProfilesNil sets the value for ClusterProfiles to be an explicit nil

### UnsetClusterProfiles
`func (o *HyperflexLocalCredentialPolicyRelationship) UnsetClusterProfiles()`

UnsetClusterProfiles ensures that no value is present for ClusterProfiles, not even an explicit nil
### GetOrganization

`func (o *HyperflexLocalCredentialPolicyRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexLocalCredentialPolicyRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexLocalCredentialPolicyRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexLocalCredentialPolicyRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


