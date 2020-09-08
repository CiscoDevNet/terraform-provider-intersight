# FabricConfigChangeDetailRelationship

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
**Changes** | Pointer to **[]string** |  | [optional] 
**ConfigChangeContext** | Pointer to [**PolicyConfigResultContext**](policy.ConfigResultContext.md) |  | [optional] 
**ConfigChangeFlag** | Pointer to **string** | Config change flag to differentiate Pending-changes and Config-drift. * &#x60;Pending-changes&#x60; - Config change flag represents changes are due to not deployed changes from Intersight. * &#x60;Drift-changes&#x60; - Config change flag represents changes are due to endpoint configuration changes. | [optional] [default to "Pending-changes"]
**Disruptions** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** | Detailed description of the config change. | [optional] 
**ModStatus** | Pointer to **string** | Modification status of the mo in this config change. * &#x60;None&#x60; - The &#39;none&#39; operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration. * &#x60;Created&#x60; - The &#39;create&#39; operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet. * &#x60;Modified&#x60; - The &#39;update&#39; operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully. * &#x60;Deleted&#x60; - The &#39;delete&#39; operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet. | [optional] [default to "None"]
**Profile** | Pointer to [**FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) |  | [optional] 

## Methods

### NewFabricConfigChangeDetailRelationship

`func NewFabricConfigChangeDetailRelationship(classId string, objectType string, ) *FabricConfigChangeDetailRelationship`

NewFabricConfigChangeDetailRelationship instantiates a new FabricConfigChangeDetailRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricConfigChangeDetailRelationshipWithDefaults

`func NewFabricConfigChangeDetailRelationshipWithDefaults() *FabricConfigChangeDetailRelationship`

NewFabricConfigChangeDetailRelationshipWithDefaults instantiates a new FabricConfigChangeDetailRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *FabricConfigChangeDetailRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *FabricConfigChangeDetailRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *FabricConfigChangeDetailRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *FabricConfigChangeDetailRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *FabricConfigChangeDetailRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *FabricConfigChangeDetailRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *FabricConfigChangeDetailRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *FabricConfigChangeDetailRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *FabricConfigChangeDetailRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *FabricConfigChangeDetailRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *FabricConfigChangeDetailRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *FabricConfigChangeDetailRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *FabricConfigChangeDetailRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *FabricConfigChangeDetailRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *FabricConfigChangeDetailRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *FabricConfigChangeDetailRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *FabricConfigChangeDetailRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *FabricConfigChangeDetailRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *FabricConfigChangeDetailRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *FabricConfigChangeDetailRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *FabricConfigChangeDetailRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *FabricConfigChangeDetailRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *FabricConfigChangeDetailRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *FabricConfigChangeDetailRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *FabricConfigChangeDetailRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *FabricConfigChangeDetailRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *FabricConfigChangeDetailRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *FabricConfigChangeDetailRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *FabricConfigChangeDetailRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *FabricConfigChangeDetailRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *FabricConfigChangeDetailRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *FabricConfigChangeDetailRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *FabricConfigChangeDetailRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *FabricConfigChangeDetailRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *FabricConfigChangeDetailRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FabricConfigChangeDetailRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FabricConfigChangeDetailRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *FabricConfigChangeDetailRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *FabricConfigChangeDetailRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *FabricConfigChangeDetailRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *FabricConfigChangeDetailRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *FabricConfigChangeDetailRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *FabricConfigChangeDetailRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *FabricConfigChangeDetailRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *FabricConfigChangeDetailRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *FabricConfigChangeDetailRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *FabricConfigChangeDetailRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *FabricConfigChangeDetailRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *FabricConfigChangeDetailRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FabricConfigChangeDetailRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FabricConfigChangeDetailRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FabricConfigChangeDetailRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *FabricConfigChangeDetailRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *FabricConfigChangeDetailRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *FabricConfigChangeDetailRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *FabricConfigChangeDetailRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *FabricConfigChangeDetailRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *FabricConfigChangeDetailRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *FabricConfigChangeDetailRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *FabricConfigChangeDetailRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *FabricConfigChangeDetailRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *FabricConfigChangeDetailRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *FabricConfigChangeDetailRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *FabricConfigChangeDetailRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetChanges

`func (o *FabricConfigChangeDetailRelationship) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *FabricConfigChangeDetailRelationship) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *FabricConfigChangeDetailRelationship) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *FabricConfigChangeDetailRelationship) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetConfigChangeContext

`func (o *FabricConfigChangeDetailRelationship) GetConfigChangeContext() PolicyConfigResultContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *FabricConfigChangeDetailRelationship) GetConfigChangeContextOk() (*PolicyConfigResultContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *FabricConfigChangeDetailRelationship) SetConfigChangeContext(v PolicyConfigResultContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *FabricConfigChangeDetailRelationship) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### GetConfigChangeFlag

`func (o *FabricConfigChangeDetailRelationship) GetConfigChangeFlag() string`

GetConfigChangeFlag returns the ConfigChangeFlag field if non-nil, zero value otherwise.

### GetConfigChangeFlagOk

`func (o *FabricConfigChangeDetailRelationship) GetConfigChangeFlagOk() (*string, bool)`

GetConfigChangeFlagOk returns a tuple with the ConfigChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeFlag

`func (o *FabricConfigChangeDetailRelationship) SetConfigChangeFlag(v string)`

SetConfigChangeFlag sets ConfigChangeFlag field to given value.

### HasConfigChangeFlag

`func (o *FabricConfigChangeDetailRelationship) HasConfigChangeFlag() bool`

HasConfigChangeFlag returns a boolean if a field has been set.

### GetDisruptions

`func (o *FabricConfigChangeDetailRelationship) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *FabricConfigChangeDetailRelationship) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *FabricConfigChangeDetailRelationship) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *FabricConfigChangeDetailRelationship) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### GetMessage

`func (o *FabricConfigChangeDetailRelationship) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FabricConfigChangeDetailRelationship) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FabricConfigChangeDetailRelationship) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FabricConfigChangeDetailRelationship) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModStatus

`func (o *FabricConfigChangeDetailRelationship) GetModStatus() string`

GetModStatus returns the ModStatus field if non-nil, zero value otherwise.

### GetModStatusOk

`func (o *FabricConfigChangeDetailRelationship) GetModStatusOk() (*string, bool)`

GetModStatusOk returns a tuple with the ModStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModStatus

`func (o *FabricConfigChangeDetailRelationship) SetModStatus(v string)`

SetModStatus sets ModStatus field to given value.

### HasModStatus

`func (o *FabricConfigChangeDetailRelationship) HasModStatus() bool`

HasModStatus returns a boolean if a field has been set.

### GetProfile

`func (o *FabricConfigChangeDetailRelationship) GetProfile() FabricSwitchProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *FabricConfigChangeDetailRelationship) GetProfileOk() (*FabricSwitchProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *FabricConfigChangeDetailRelationship) SetProfile(v FabricSwitchProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *FabricConfigChangeDetailRelationship) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


