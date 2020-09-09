# SdwanRouterNodeRelationship

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
**DeviceTemplate** | Pointer to **string** | Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration. | [optional] 
**Name** | Pointer to **string** | Name of the router node object. | [optional] 
**NetworkConfiguration** | Pointer to [**[]SdwanNetworkConfigurationType**](sdwan.NetworkConfigurationType.md) |  | [optional] 
**TemplateInputs** | Pointer to [**[]SdwanTemplateInputsType**](sdwan.TemplateInputsType.md) |  | [optional] 
**Uuid** | Pointer to **string** | Uniquely identifies the router by its chassis number. | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profile** | Pointer to [**SdwanProfileRelationship**](sdwan.Profile.Relationship.md) |  | [optional] 
**ServerNode** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewSdwanRouterNodeRelationship

`func NewSdwanRouterNodeRelationship(classId string, objectType string, ) *SdwanRouterNodeRelationship`

NewSdwanRouterNodeRelationship instantiates a new SdwanRouterNodeRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSdwanRouterNodeRelationshipWithDefaults

`func NewSdwanRouterNodeRelationshipWithDefaults() *SdwanRouterNodeRelationship`

NewSdwanRouterNodeRelationshipWithDefaults instantiates a new SdwanRouterNodeRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *SdwanRouterNodeRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *SdwanRouterNodeRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *SdwanRouterNodeRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *SdwanRouterNodeRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *SdwanRouterNodeRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *SdwanRouterNodeRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *SdwanRouterNodeRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *SdwanRouterNodeRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *SdwanRouterNodeRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *SdwanRouterNodeRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *SdwanRouterNodeRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *SdwanRouterNodeRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *SdwanRouterNodeRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *SdwanRouterNodeRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *SdwanRouterNodeRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *SdwanRouterNodeRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *SdwanRouterNodeRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *SdwanRouterNodeRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *SdwanRouterNodeRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *SdwanRouterNodeRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *SdwanRouterNodeRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *SdwanRouterNodeRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *SdwanRouterNodeRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *SdwanRouterNodeRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *SdwanRouterNodeRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *SdwanRouterNodeRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *SdwanRouterNodeRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *SdwanRouterNodeRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *SdwanRouterNodeRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *SdwanRouterNodeRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *SdwanRouterNodeRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *SdwanRouterNodeRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *SdwanRouterNodeRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *SdwanRouterNodeRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *SdwanRouterNodeRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SdwanRouterNodeRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SdwanRouterNodeRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SdwanRouterNodeRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *SdwanRouterNodeRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *SdwanRouterNodeRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *SdwanRouterNodeRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *SdwanRouterNodeRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *SdwanRouterNodeRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *SdwanRouterNodeRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *SdwanRouterNodeRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *SdwanRouterNodeRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *SdwanRouterNodeRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *SdwanRouterNodeRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *SdwanRouterNodeRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *SdwanRouterNodeRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *SdwanRouterNodeRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *SdwanRouterNodeRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *SdwanRouterNodeRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *SdwanRouterNodeRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *SdwanRouterNodeRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *SdwanRouterNodeRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *SdwanRouterNodeRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *SdwanRouterNodeRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *SdwanRouterNodeRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *SdwanRouterNodeRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *SdwanRouterNodeRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *SdwanRouterNodeRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *SdwanRouterNodeRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *SdwanRouterNodeRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDeviceTemplate

`func (o *SdwanRouterNodeRelationship) GetDeviceTemplate() string`

GetDeviceTemplate returns the DeviceTemplate field if non-nil, zero value otherwise.

### GetDeviceTemplateOk

`func (o *SdwanRouterNodeRelationship) GetDeviceTemplateOk() (*string, bool)`

GetDeviceTemplateOk returns a tuple with the DeviceTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceTemplate

`func (o *SdwanRouterNodeRelationship) SetDeviceTemplate(v string)`

SetDeviceTemplate sets DeviceTemplate field to given value.

### HasDeviceTemplate

`func (o *SdwanRouterNodeRelationship) HasDeviceTemplate() bool`

HasDeviceTemplate returns a boolean if a field has been set.

### GetName

`func (o *SdwanRouterNodeRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SdwanRouterNodeRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SdwanRouterNodeRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SdwanRouterNodeRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkConfiguration

`func (o *SdwanRouterNodeRelationship) GetNetworkConfiguration() []SdwanNetworkConfigurationType`

GetNetworkConfiguration returns the NetworkConfiguration field if non-nil, zero value otherwise.

### GetNetworkConfigurationOk

`func (o *SdwanRouterNodeRelationship) GetNetworkConfigurationOk() (*[]SdwanNetworkConfigurationType, bool)`

GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkConfiguration

`func (o *SdwanRouterNodeRelationship) SetNetworkConfiguration(v []SdwanNetworkConfigurationType)`

SetNetworkConfiguration sets NetworkConfiguration field to given value.

### HasNetworkConfiguration

`func (o *SdwanRouterNodeRelationship) HasNetworkConfiguration() bool`

HasNetworkConfiguration returns a boolean if a field has been set.

### GetTemplateInputs

`func (o *SdwanRouterNodeRelationship) GetTemplateInputs() []SdwanTemplateInputsType`

GetTemplateInputs returns the TemplateInputs field if non-nil, zero value otherwise.

### GetTemplateInputsOk

`func (o *SdwanRouterNodeRelationship) GetTemplateInputsOk() (*[]SdwanTemplateInputsType, bool)`

GetTemplateInputsOk returns a tuple with the TemplateInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateInputs

`func (o *SdwanRouterNodeRelationship) SetTemplateInputs(v []SdwanTemplateInputsType)`

SetTemplateInputs sets TemplateInputs field to given value.

### HasTemplateInputs

`func (o *SdwanRouterNodeRelationship) HasTemplateInputs() bool`

HasTemplateInputs returns a boolean if a field has been set.

### GetUuid

`func (o *SdwanRouterNodeRelationship) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SdwanRouterNodeRelationship) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SdwanRouterNodeRelationship) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SdwanRouterNodeRelationship) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetOrganization

`func (o *SdwanRouterNodeRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SdwanRouterNodeRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SdwanRouterNodeRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SdwanRouterNodeRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfile

`func (o *SdwanRouterNodeRelationship) GetProfile() SdwanProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *SdwanRouterNodeRelationship) GetProfileOk() (*SdwanProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *SdwanRouterNodeRelationship) SetProfile(v SdwanProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *SdwanRouterNodeRelationship) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetServerNode

`func (o *SdwanRouterNodeRelationship) GetServerNode() AssetDeviceRegistrationRelationship`

GetServerNode returns the ServerNode field if non-nil, zero value otherwise.

### GetServerNodeOk

`func (o *SdwanRouterNodeRelationship) GetServerNodeOk() (*AssetDeviceRegistrationRelationship, bool)`

GetServerNodeOk returns a tuple with the ServerNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNode

`func (o *SdwanRouterNodeRelationship) SetServerNode(v AssetDeviceRegistrationRelationship)`

SetServerNode sets ServerNode field to given value.

### HasServerNode

`func (o *SdwanRouterNodeRelationship) HasServerNode() bool`

HasServerNode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


