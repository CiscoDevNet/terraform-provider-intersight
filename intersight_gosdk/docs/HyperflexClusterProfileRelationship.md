# HyperflexClusterProfileRelationship

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
**Action** | Pointer to **string** | User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign. | [optional] 
**ConfigContext** | Pointer to [**PolicyConfigContext**](policy.ConfigContext.md) |  | [optional] 
**DataIpAddress** | Pointer to **string** | The storage data IP address for the HyperFlex cluster. | [optional] 
**HypervisorType** | Pointer to **string** | The hypervisor type for the HyperFlex cluster. * &#x60;ESXi&#x60; - ESXi hypervisor as specified by the user. * &#x60;HYPERV&#x60; - Hyperv hypervisor as specified by the user. * &#x60;KVM&#x60; - KVM hypervisor as specified by the user. | [optional] [default to "ESXi"]
**MacAddressPrefix** | Pointer to **string** | The MAC address prefix in the form of 00:25:B5:XX. | [optional] 
**MgmtIpAddress** | Pointer to **string** | The management IP address for the HyperFlex cluster. | [optional] 
**MgmtPlatform** | Pointer to **string** | The management platform for the HyperFlex cluster. * &#x60;FI&#x60; - The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect. * &#x60;EDGE&#x60; - The host servers used in the cluster deployment are standalone severs. | [optional] [default to "FI"]
**Replication** | Pointer to **int64** | The number of copies of each data block written. | [optional] 
**StorageDataVlan** | Pointer to [**HyperflexNamedVlan**](hyperflex.NamedVlan.md) |  | [optional] 
**WwxnPrefix** | Pointer to **string** | The WWxN prefix in the form of 20:00:00:25:B5:XX. | [optional] 
**AssociatedCluster** | Pointer to [**HyperflexClusterRelationship**](hyperflex.Cluster.Relationship.md) |  | [optional] 
**AutoSupport** | Pointer to [**HyperflexAutoSupportPolicyRelationship**](hyperflex.AutoSupportPolicy.Relationship.md) |  | [optional] 
**ClusterNetwork** | Pointer to [**HyperflexClusterNetworkPolicyRelationship**](hyperflex.ClusterNetworkPolicy.Relationship.md) |  | [optional] 
**ClusterStorage** | Pointer to [**HyperflexClusterStoragePolicyRelationship**](hyperflex.ClusterStoragePolicy.Relationship.md) |  | [optional] 
**ConfigResult** | Pointer to [**HyperflexConfigResultRelationship**](hyperflex.ConfigResult.Relationship.md) |  | [optional] 
**ExtFcStorage** | Pointer to [**HyperflexExtFcStoragePolicyRelationship**](hyperflex.ExtFcStoragePolicy.Relationship.md) |  | [optional] 
**ExtIscsiStorage** | Pointer to [**HyperflexExtIscsiStoragePolicyRelationship**](hyperflex.ExtIscsiStoragePolicy.Relationship.md) |  | [optional] 
**LocalCredential** | Pointer to [**HyperflexLocalCredentialPolicyRelationship**](hyperflex.LocalCredentialPolicy.Relationship.md) |  | [optional] 
**NodeConfig** | Pointer to [**HyperflexNodeConfigPolicyRelationship**](hyperflex.NodeConfigPolicy.Relationship.md) |  | [optional] 
**NodeProfileConfig** | Pointer to [**[]HyperflexNodeProfileRelationship**](hyperflex.NodeProfile.Relationship.md) | An array of relationships to hyperflexNodeProfile resources. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**ProxySetting** | Pointer to [**HyperflexProxySettingPolicyRelationship**](hyperflex.ProxySettingPolicy.Relationship.md) |  | [optional] 
**RunningWorkflows** | Pointer to [**[]WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) | An array of relationships to workflowWorkflowInfo resources. | [optional] [readonly] 
**SoftwareVersion** | Pointer to [**HyperflexSoftwareVersionPolicyRelationship**](hyperflex.SoftwareVersionPolicy.Relationship.md) |  | [optional] 
**SysConfig** | Pointer to [**HyperflexSysConfigPolicyRelationship**](hyperflex.SysConfigPolicy.Relationship.md) |  | [optional] 
**UcsmConfig** | Pointer to [**HyperflexUcsmConfigPolicyRelationship**](hyperflex.UcsmConfigPolicy.Relationship.md) |  | [optional] 
**VcenterConfig** | Pointer to [**HyperflexVcenterConfigPolicyRelationship**](hyperflex.VcenterConfigPolicy.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexClusterProfileRelationship

`func NewHyperflexClusterProfileRelationship(classId string, objectType string, ) *HyperflexClusterProfileRelationship`

NewHyperflexClusterProfileRelationship instantiates a new HyperflexClusterProfileRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexClusterProfileRelationshipWithDefaults

`func NewHyperflexClusterProfileRelationshipWithDefaults() *HyperflexClusterProfileRelationship`

NewHyperflexClusterProfileRelationshipWithDefaults instantiates a new HyperflexClusterProfileRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *HyperflexClusterProfileRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *HyperflexClusterProfileRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *HyperflexClusterProfileRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *HyperflexClusterProfileRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *HyperflexClusterProfileRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *HyperflexClusterProfileRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *HyperflexClusterProfileRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *HyperflexClusterProfileRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *HyperflexClusterProfileRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *HyperflexClusterProfileRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *HyperflexClusterProfileRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *HyperflexClusterProfileRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *HyperflexClusterProfileRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *HyperflexClusterProfileRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *HyperflexClusterProfileRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *HyperflexClusterProfileRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *HyperflexClusterProfileRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *HyperflexClusterProfileRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *HyperflexClusterProfileRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *HyperflexClusterProfileRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *HyperflexClusterProfileRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *HyperflexClusterProfileRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *HyperflexClusterProfileRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *HyperflexClusterProfileRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *HyperflexClusterProfileRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *HyperflexClusterProfileRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *HyperflexClusterProfileRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *HyperflexClusterProfileRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *HyperflexClusterProfileRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *HyperflexClusterProfileRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *HyperflexClusterProfileRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *HyperflexClusterProfileRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *HyperflexClusterProfileRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *HyperflexClusterProfileRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *HyperflexClusterProfileRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *HyperflexClusterProfileRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *HyperflexClusterProfileRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *HyperflexClusterProfileRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *HyperflexClusterProfileRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *HyperflexClusterProfileRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *HyperflexClusterProfileRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *HyperflexClusterProfileRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *HyperflexClusterProfileRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *HyperflexClusterProfileRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *HyperflexClusterProfileRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *HyperflexClusterProfileRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *HyperflexClusterProfileRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *HyperflexClusterProfileRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *HyperflexClusterProfileRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *HyperflexClusterProfileRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *HyperflexClusterProfileRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *HyperflexClusterProfileRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *HyperflexClusterProfileRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *HyperflexClusterProfileRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *HyperflexClusterProfileRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *HyperflexClusterProfileRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *HyperflexClusterProfileRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *HyperflexClusterProfileRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *HyperflexClusterProfileRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *HyperflexClusterProfileRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *HyperflexClusterProfileRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *HyperflexClusterProfileRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *HyperflexClusterProfileRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *HyperflexClusterProfileRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDescription

`func (o *HyperflexClusterProfileRelationship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *HyperflexClusterProfileRelationship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *HyperflexClusterProfileRelationship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *HyperflexClusterProfileRelationship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *HyperflexClusterProfileRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexClusterProfileRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexClusterProfileRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexClusterProfileRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *HyperflexClusterProfileRelationship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexClusterProfileRelationship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexClusterProfileRelationship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexClusterProfileRelationship) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSrcTemplate

`func (o *HyperflexClusterProfileRelationship) GetSrcTemplate() PolicyAbstractProfileRelationship`

GetSrcTemplate returns the SrcTemplate field if non-nil, zero value otherwise.

### GetSrcTemplateOk

`func (o *HyperflexClusterProfileRelationship) GetSrcTemplateOk() (*PolicyAbstractProfileRelationship, bool)`

GetSrcTemplateOk returns a tuple with the SrcTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcTemplate

`func (o *HyperflexClusterProfileRelationship) SetSrcTemplate(v PolicyAbstractProfileRelationship)`

SetSrcTemplate sets SrcTemplate field to given value.

### HasSrcTemplate

`func (o *HyperflexClusterProfileRelationship) HasSrcTemplate() bool`

HasSrcTemplate returns a boolean if a field has been set.

### GetAction

`func (o *HyperflexClusterProfileRelationship) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *HyperflexClusterProfileRelationship) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *HyperflexClusterProfileRelationship) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *HyperflexClusterProfileRelationship) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetConfigContext

`func (o *HyperflexClusterProfileRelationship) GetConfigContext() PolicyConfigContext`

GetConfigContext returns the ConfigContext field if non-nil, zero value otherwise.

### GetConfigContextOk

`func (o *HyperflexClusterProfileRelationship) GetConfigContextOk() (*PolicyConfigContext, bool)`

GetConfigContextOk returns a tuple with the ConfigContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigContext

`func (o *HyperflexClusterProfileRelationship) SetConfigContext(v PolicyConfigContext)`

SetConfigContext sets ConfigContext field to given value.

### HasConfigContext

`func (o *HyperflexClusterProfileRelationship) HasConfigContext() bool`

HasConfigContext returns a boolean if a field has been set.

### GetDataIpAddress

`func (o *HyperflexClusterProfileRelationship) GetDataIpAddress() string`

GetDataIpAddress returns the DataIpAddress field if non-nil, zero value otherwise.

### GetDataIpAddressOk

`func (o *HyperflexClusterProfileRelationship) GetDataIpAddressOk() (*string, bool)`

GetDataIpAddressOk returns a tuple with the DataIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataIpAddress

`func (o *HyperflexClusterProfileRelationship) SetDataIpAddress(v string)`

SetDataIpAddress sets DataIpAddress field to given value.

### HasDataIpAddress

`func (o *HyperflexClusterProfileRelationship) HasDataIpAddress() bool`

HasDataIpAddress returns a boolean if a field has been set.

### GetHypervisorType

`func (o *HyperflexClusterProfileRelationship) GetHypervisorType() string`

GetHypervisorType returns the HypervisorType field if non-nil, zero value otherwise.

### GetHypervisorTypeOk

`func (o *HyperflexClusterProfileRelationship) GetHypervisorTypeOk() (*string, bool)`

GetHypervisorTypeOk returns a tuple with the HypervisorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorType

`func (o *HyperflexClusterProfileRelationship) SetHypervisorType(v string)`

SetHypervisorType sets HypervisorType field to given value.

### HasHypervisorType

`func (o *HyperflexClusterProfileRelationship) HasHypervisorType() bool`

HasHypervisorType returns a boolean if a field has been set.

### GetMacAddressPrefix

`func (o *HyperflexClusterProfileRelationship) GetMacAddressPrefix() string`

GetMacAddressPrefix returns the MacAddressPrefix field if non-nil, zero value otherwise.

### GetMacAddressPrefixOk

`func (o *HyperflexClusterProfileRelationship) GetMacAddressPrefixOk() (*string, bool)`

GetMacAddressPrefixOk returns a tuple with the MacAddressPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddressPrefix

`func (o *HyperflexClusterProfileRelationship) SetMacAddressPrefix(v string)`

SetMacAddressPrefix sets MacAddressPrefix field to given value.

### HasMacAddressPrefix

`func (o *HyperflexClusterProfileRelationship) HasMacAddressPrefix() bool`

HasMacAddressPrefix returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *HyperflexClusterProfileRelationship) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *HyperflexClusterProfileRelationship) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *HyperflexClusterProfileRelationship) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *HyperflexClusterProfileRelationship) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *HyperflexClusterProfileRelationship) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *HyperflexClusterProfileRelationship) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *HyperflexClusterProfileRelationship) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *HyperflexClusterProfileRelationship) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetReplication

`func (o *HyperflexClusterProfileRelationship) GetReplication() int64`

GetReplication returns the Replication field if non-nil, zero value otherwise.

### GetReplicationOk

`func (o *HyperflexClusterProfileRelationship) GetReplicationOk() (*int64, bool)`

GetReplicationOk returns a tuple with the Replication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplication

`func (o *HyperflexClusterProfileRelationship) SetReplication(v int64)`

SetReplication sets Replication field to given value.

### HasReplication

`func (o *HyperflexClusterProfileRelationship) HasReplication() bool`

HasReplication returns a boolean if a field has been set.

### GetStorageDataVlan

`func (o *HyperflexClusterProfileRelationship) GetStorageDataVlan() HyperflexNamedVlan`

GetStorageDataVlan returns the StorageDataVlan field if non-nil, zero value otherwise.

### GetStorageDataVlanOk

`func (o *HyperflexClusterProfileRelationship) GetStorageDataVlanOk() (*HyperflexNamedVlan, bool)`

GetStorageDataVlanOk returns a tuple with the StorageDataVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageDataVlan

`func (o *HyperflexClusterProfileRelationship) SetStorageDataVlan(v HyperflexNamedVlan)`

SetStorageDataVlan sets StorageDataVlan field to given value.

### HasStorageDataVlan

`func (o *HyperflexClusterProfileRelationship) HasStorageDataVlan() bool`

HasStorageDataVlan returns a boolean if a field has been set.

### GetWwxnPrefix

`func (o *HyperflexClusterProfileRelationship) GetWwxnPrefix() string`

GetWwxnPrefix returns the WwxnPrefix field if non-nil, zero value otherwise.

### GetWwxnPrefixOk

`func (o *HyperflexClusterProfileRelationship) GetWwxnPrefixOk() (*string, bool)`

GetWwxnPrefixOk returns a tuple with the WwxnPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWwxnPrefix

`func (o *HyperflexClusterProfileRelationship) SetWwxnPrefix(v string)`

SetWwxnPrefix sets WwxnPrefix field to given value.

### HasWwxnPrefix

`func (o *HyperflexClusterProfileRelationship) HasWwxnPrefix() bool`

HasWwxnPrefix returns a boolean if a field has been set.

### GetAssociatedCluster

`func (o *HyperflexClusterProfileRelationship) GetAssociatedCluster() HyperflexClusterRelationship`

GetAssociatedCluster returns the AssociatedCluster field if non-nil, zero value otherwise.

### GetAssociatedClusterOk

`func (o *HyperflexClusterProfileRelationship) GetAssociatedClusterOk() (*HyperflexClusterRelationship, bool)`

GetAssociatedClusterOk returns a tuple with the AssociatedCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedCluster

`func (o *HyperflexClusterProfileRelationship) SetAssociatedCluster(v HyperflexClusterRelationship)`

SetAssociatedCluster sets AssociatedCluster field to given value.

### HasAssociatedCluster

`func (o *HyperflexClusterProfileRelationship) HasAssociatedCluster() bool`

HasAssociatedCluster returns a boolean if a field has been set.

### GetAutoSupport

`func (o *HyperflexClusterProfileRelationship) GetAutoSupport() HyperflexAutoSupportPolicyRelationship`

GetAutoSupport returns the AutoSupport field if non-nil, zero value otherwise.

### GetAutoSupportOk

`func (o *HyperflexClusterProfileRelationship) GetAutoSupportOk() (*HyperflexAutoSupportPolicyRelationship, bool)`

GetAutoSupportOk returns a tuple with the AutoSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoSupport

`func (o *HyperflexClusterProfileRelationship) SetAutoSupport(v HyperflexAutoSupportPolicyRelationship)`

SetAutoSupport sets AutoSupport field to given value.

### HasAutoSupport

`func (o *HyperflexClusterProfileRelationship) HasAutoSupport() bool`

HasAutoSupport returns a boolean if a field has been set.

### GetClusterNetwork

`func (o *HyperflexClusterProfileRelationship) GetClusterNetwork() HyperflexClusterNetworkPolicyRelationship`

GetClusterNetwork returns the ClusterNetwork field if non-nil, zero value otherwise.

### GetClusterNetworkOk

`func (o *HyperflexClusterProfileRelationship) GetClusterNetworkOk() (*HyperflexClusterNetworkPolicyRelationship, bool)`

GetClusterNetworkOk returns a tuple with the ClusterNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterNetwork

`func (o *HyperflexClusterProfileRelationship) SetClusterNetwork(v HyperflexClusterNetworkPolicyRelationship)`

SetClusterNetwork sets ClusterNetwork field to given value.

### HasClusterNetwork

`func (o *HyperflexClusterProfileRelationship) HasClusterNetwork() bool`

HasClusterNetwork returns a boolean if a field has been set.

### GetClusterStorage

`func (o *HyperflexClusterProfileRelationship) GetClusterStorage() HyperflexClusterStoragePolicyRelationship`

GetClusterStorage returns the ClusterStorage field if non-nil, zero value otherwise.

### GetClusterStorageOk

`func (o *HyperflexClusterProfileRelationship) GetClusterStorageOk() (*HyperflexClusterStoragePolicyRelationship, bool)`

GetClusterStorageOk returns a tuple with the ClusterStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterStorage

`func (o *HyperflexClusterProfileRelationship) SetClusterStorage(v HyperflexClusterStoragePolicyRelationship)`

SetClusterStorage sets ClusterStorage field to given value.

### HasClusterStorage

`func (o *HyperflexClusterProfileRelationship) HasClusterStorage() bool`

HasClusterStorage returns a boolean if a field has been set.

### GetConfigResult

`func (o *HyperflexClusterProfileRelationship) GetConfigResult() HyperflexConfigResultRelationship`

GetConfigResult returns the ConfigResult field if non-nil, zero value otherwise.

### GetConfigResultOk

`func (o *HyperflexClusterProfileRelationship) GetConfigResultOk() (*HyperflexConfigResultRelationship, bool)`

GetConfigResultOk returns a tuple with the ConfigResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigResult

`func (o *HyperflexClusterProfileRelationship) SetConfigResult(v HyperflexConfigResultRelationship)`

SetConfigResult sets ConfigResult field to given value.

### HasConfigResult

`func (o *HyperflexClusterProfileRelationship) HasConfigResult() bool`

HasConfigResult returns a boolean if a field has been set.

### GetExtFcStorage

`func (o *HyperflexClusterProfileRelationship) GetExtFcStorage() HyperflexExtFcStoragePolicyRelationship`

GetExtFcStorage returns the ExtFcStorage field if non-nil, zero value otherwise.

### GetExtFcStorageOk

`func (o *HyperflexClusterProfileRelationship) GetExtFcStorageOk() (*HyperflexExtFcStoragePolicyRelationship, bool)`

GetExtFcStorageOk returns a tuple with the ExtFcStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtFcStorage

`func (o *HyperflexClusterProfileRelationship) SetExtFcStorage(v HyperflexExtFcStoragePolicyRelationship)`

SetExtFcStorage sets ExtFcStorage field to given value.

### HasExtFcStorage

`func (o *HyperflexClusterProfileRelationship) HasExtFcStorage() bool`

HasExtFcStorage returns a boolean if a field has been set.

### GetExtIscsiStorage

`func (o *HyperflexClusterProfileRelationship) GetExtIscsiStorage() HyperflexExtIscsiStoragePolicyRelationship`

GetExtIscsiStorage returns the ExtIscsiStorage field if non-nil, zero value otherwise.

### GetExtIscsiStorageOk

`func (o *HyperflexClusterProfileRelationship) GetExtIscsiStorageOk() (*HyperflexExtIscsiStoragePolicyRelationship, bool)`

GetExtIscsiStorageOk returns a tuple with the ExtIscsiStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtIscsiStorage

`func (o *HyperflexClusterProfileRelationship) SetExtIscsiStorage(v HyperflexExtIscsiStoragePolicyRelationship)`

SetExtIscsiStorage sets ExtIscsiStorage field to given value.

### HasExtIscsiStorage

`func (o *HyperflexClusterProfileRelationship) HasExtIscsiStorage() bool`

HasExtIscsiStorage returns a boolean if a field has been set.

### GetLocalCredential

`func (o *HyperflexClusterProfileRelationship) GetLocalCredential() HyperflexLocalCredentialPolicyRelationship`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *HyperflexClusterProfileRelationship) GetLocalCredentialOk() (*HyperflexLocalCredentialPolicyRelationship, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *HyperflexClusterProfileRelationship) SetLocalCredential(v HyperflexLocalCredentialPolicyRelationship)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *HyperflexClusterProfileRelationship) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### GetNodeConfig

`func (o *HyperflexClusterProfileRelationship) GetNodeConfig() HyperflexNodeConfigPolicyRelationship`

GetNodeConfig returns the NodeConfig field if non-nil, zero value otherwise.

### GetNodeConfigOk

`func (o *HyperflexClusterProfileRelationship) GetNodeConfigOk() (*HyperflexNodeConfigPolicyRelationship, bool)`

GetNodeConfigOk returns a tuple with the NodeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeConfig

`func (o *HyperflexClusterProfileRelationship) SetNodeConfig(v HyperflexNodeConfigPolicyRelationship)`

SetNodeConfig sets NodeConfig field to given value.

### HasNodeConfig

`func (o *HyperflexClusterProfileRelationship) HasNodeConfig() bool`

HasNodeConfig returns a boolean if a field has been set.

### GetNodeProfileConfig

`func (o *HyperflexClusterProfileRelationship) GetNodeProfileConfig() []HyperflexNodeProfileRelationship`

GetNodeProfileConfig returns the NodeProfileConfig field if non-nil, zero value otherwise.

### GetNodeProfileConfigOk

`func (o *HyperflexClusterProfileRelationship) GetNodeProfileConfigOk() (*[]HyperflexNodeProfileRelationship, bool)`

GetNodeProfileConfigOk returns a tuple with the NodeProfileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeProfileConfig

`func (o *HyperflexClusterProfileRelationship) SetNodeProfileConfig(v []HyperflexNodeProfileRelationship)`

SetNodeProfileConfig sets NodeProfileConfig field to given value.

### HasNodeProfileConfig

`func (o *HyperflexClusterProfileRelationship) HasNodeProfileConfig() bool`

HasNodeProfileConfig returns a boolean if a field has been set.

### SetNodeProfileConfigNil

`func (o *HyperflexClusterProfileRelationship) SetNodeProfileConfigNil(b bool)`

 SetNodeProfileConfigNil sets the value for NodeProfileConfig to be an explicit nil

### UnsetNodeProfileConfig
`func (o *HyperflexClusterProfileRelationship) UnsetNodeProfileConfig()`

UnsetNodeProfileConfig ensures that no value is present for NodeProfileConfig, not even an explicit nil
### GetOrganization

`func (o *HyperflexClusterProfileRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *HyperflexClusterProfileRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *HyperflexClusterProfileRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *HyperflexClusterProfileRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProxySetting

`func (o *HyperflexClusterProfileRelationship) GetProxySetting() HyperflexProxySettingPolicyRelationship`

GetProxySetting returns the ProxySetting field if non-nil, zero value otherwise.

### GetProxySettingOk

`func (o *HyperflexClusterProfileRelationship) GetProxySettingOk() (*HyperflexProxySettingPolicyRelationship, bool)`

GetProxySettingOk returns a tuple with the ProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxySetting

`func (o *HyperflexClusterProfileRelationship) SetProxySetting(v HyperflexProxySettingPolicyRelationship)`

SetProxySetting sets ProxySetting field to given value.

### HasProxySetting

`func (o *HyperflexClusterProfileRelationship) HasProxySetting() bool`

HasProxySetting returns a boolean if a field has been set.

### GetRunningWorkflows

`func (o *HyperflexClusterProfileRelationship) GetRunningWorkflows() []WorkflowWorkflowInfoRelationship`

GetRunningWorkflows returns the RunningWorkflows field if non-nil, zero value otherwise.

### GetRunningWorkflowsOk

`func (o *HyperflexClusterProfileRelationship) GetRunningWorkflowsOk() (*[]WorkflowWorkflowInfoRelationship, bool)`

GetRunningWorkflowsOk returns a tuple with the RunningWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningWorkflows

`func (o *HyperflexClusterProfileRelationship) SetRunningWorkflows(v []WorkflowWorkflowInfoRelationship)`

SetRunningWorkflows sets RunningWorkflows field to given value.

### HasRunningWorkflows

`func (o *HyperflexClusterProfileRelationship) HasRunningWorkflows() bool`

HasRunningWorkflows returns a boolean if a field has been set.

### SetRunningWorkflowsNil

`func (o *HyperflexClusterProfileRelationship) SetRunningWorkflowsNil(b bool)`

 SetRunningWorkflowsNil sets the value for RunningWorkflows to be an explicit nil

### UnsetRunningWorkflows
`func (o *HyperflexClusterProfileRelationship) UnsetRunningWorkflows()`

UnsetRunningWorkflows ensures that no value is present for RunningWorkflows, not even an explicit nil
### GetSoftwareVersion

`func (o *HyperflexClusterProfileRelationship) GetSoftwareVersion() HyperflexSoftwareVersionPolicyRelationship`

GetSoftwareVersion returns the SoftwareVersion field if non-nil, zero value otherwise.

### GetSoftwareVersionOk

`func (o *HyperflexClusterProfileRelationship) GetSoftwareVersionOk() (*HyperflexSoftwareVersionPolicyRelationship, bool)`

GetSoftwareVersionOk returns a tuple with the SoftwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftwareVersion

`func (o *HyperflexClusterProfileRelationship) SetSoftwareVersion(v HyperflexSoftwareVersionPolicyRelationship)`

SetSoftwareVersion sets SoftwareVersion field to given value.

### HasSoftwareVersion

`func (o *HyperflexClusterProfileRelationship) HasSoftwareVersion() bool`

HasSoftwareVersion returns a boolean if a field has been set.

### GetSysConfig

`func (o *HyperflexClusterProfileRelationship) GetSysConfig() HyperflexSysConfigPolicyRelationship`

GetSysConfig returns the SysConfig field if non-nil, zero value otherwise.

### GetSysConfigOk

`func (o *HyperflexClusterProfileRelationship) GetSysConfigOk() (*HyperflexSysConfigPolicyRelationship, bool)`

GetSysConfigOk returns a tuple with the SysConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysConfig

`func (o *HyperflexClusterProfileRelationship) SetSysConfig(v HyperflexSysConfigPolicyRelationship)`

SetSysConfig sets SysConfig field to given value.

### HasSysConfig

`func (o *HyperflexClusterProfileRelationship) HasSysConfig() bool`

HasSysConfig returns a boolean if a field has been set.

### GetUcsmConfig

`func (o *HyperflexClusterProfileRelationship) GetUcsmConfig() HyperflexUcsmConfigPolicyRelationship`

GetUcsmConfig returns the UcsmConfig field if non-nil, zero value otherwise.

### GetUcsmConfigOk

`func (o *HyperflexClusterProfileRelationship) GetUcsmConfigOk() (*HyperflexUcsmConfigPolicyRelationship, bool)`

GetUcsmConfigOk returns a tuple with the UcsmConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUcsmConfig

`func (o *HyperflexClusterProfileRelationship) SetUcsmConfig(v HyperflexUcsmConfigPolicyRelationship)`

SetUcsmConfig sets UcsmConfig field to given value.

### HasUcsmConfig

`func (o *HyperflexClusterProfileRelationship) HasUcsmConfig() bool`

HasUcsmConfig returns a boolean if a field has been set.

### GetVcenterConfig

`func (o *HyperflexClusterProfileRelationship) GetVcenterConfig() HyperflexVcenterConfigPolicyRelationship`

GetVcenterConfig returns the VcenterConfig field if non-nil, zero value otherwise.

### GetVcenterConfigOk

`func (o *HyperflexClusterProfileRelationship) GetVcenterConfigOk() (*HyperflexVcenterConfigPolicyRelationship, bool)`

GetVcenterConfigOk returns a tuple with the VcenterConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterConfig

`func (o *HyperflexClusterProfileRelationship) SetVcenterConfig(v HyperflexVcenterConfigPolicyRelationship)`

SetVcenterConfig sets VcenterConfig field to given value.

### HasVcenterConfig

`func (o *HyperflexClusterProfileRelationship) HasVcenterConfig() bool`

HasVcenterConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


