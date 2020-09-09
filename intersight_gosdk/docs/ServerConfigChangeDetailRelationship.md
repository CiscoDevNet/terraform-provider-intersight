# ServerConfigChangeDetailRelationship

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
**Profile** | Pointer to [**ServerProfileRelationship**](server.Profile.Relationship.md) |  | [optional] 

## Methods

### NewServerConfigChangeDetailRelationship

`func NewServerConfigChangeDetailRelationship(classId string, objectType string, ) *ServerConfigChangeDetailRelationship`

NewServerConfigChangeDetailRelationship instantiates a new ServerConfigChangeDetailRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigChangeDetailRelationshipWithDefaults

`func NewServerConfigChangeDetailRelationshipWithDefaults() *ServerConfigChangeDetailRelationship`

NewServerConfigChangeDetailRelationshipWithDefaults instantiates a new ServerConfigChangeDetailRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ServerConfigChangeDetailRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ServerConfigChangeDetailRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ServerConfigChangeDetailRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ServerConfigChangeDetailRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ServerConfigChangeDetailRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ServerConfigChangeDetailRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ServerConfigChangeDetailRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ServerConfigChangeDetailRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ServerConfigChangeDetailRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ServerConfigChangeDetailRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ServerConfigChangeDetailRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ServerConfigChangeDetailRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ServerConfigChangeDetailRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ServerConfigChangeDetailRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ServerConfigChangeDetailRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ServerConfigChangeDetailRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ServerConfigChangeDetailRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ServerConfigChangeDetailRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ServerConfigChangeDetailRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ServerConfigChangeDetailRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ServerConfigChangeDetailRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ServerConfigChangeDetailRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ServerConfigChangeDetailRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ServerConfigChangeDetailRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ServerConfigChangeDetailRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ServerConfigChangeDetailRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ServerConfigChangeDetailRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ServerConfigChangeDetailRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ServerConfigChangeDetailRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ServerConfigChangeDetailRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ServerConfigChangeDetailRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ServerConfigChangeDetailRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ServerConfigChangeDetailRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ServerConfigChangeDetailRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ServerConfigChangeDetailRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ServerConfigChangeDetailRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ServerConfigChangeDetailRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ServerConfigChangeDetailRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ServerConfigChangeDetailRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ServerConfigChangeDetailRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ServerConfigChangeDetailRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ServerConfigChangeDetailRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ServerConfigChangeDetailRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ServerConfigChangeDetailRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ServerConfigChangeDetailRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ServerConfigChangeDetailRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ServerConfigChangeDetailRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ServerConfigChangeDetailRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ServerConfigChangeDetailRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ServerConfigChangeDetailRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ServerConfigChangeDetailRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ServerConfigChangeDetailRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ServerConfigChangeDetailRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ServerConfigChangeDetailRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ServerConfigChangeDetailRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ServerConfigChangeDetailRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ServerConfigChangeDetailRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ServerConfigChangeDetailRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ServerConfigChangeDetailRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ServerConfigChangeDetailRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ServerConfigChangeDetailRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ServerConfigChangeDetailRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ServerConfigChangeDetailRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ServerConfigChangeDetailRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetChanges

`func (o *ServerConfigChangeDetailRelationship) GetChanges() []string`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *ServerConfigChangeDetailRelationship) GetChangesOk() (*[]string, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *ServerConfigChangeDetailRelationship) SetChanges(v []string)`

SetChanges sets Changes field to given value.

### HasChanges

`func (o *ServerConfigChangeDetailRelationship) HasChanges() bool`

HasChanges returns a boolean if a field has been set.

### GetConfigChangeContext

`func (o *ServerConfigChangeDetailRelationship) GetConfigChangeContext() PolicyConfigResultContext`

GetConfigChangeContext returns the ConfigChangeContext field if non-nil, zero value otherwise.

### GetConfigChangeContextOk

`func (o *ServerConfigChangeDetailRelationship) GetConfigChangeContextOk() (*PolicyConfigResultContext, bool)`

GetConfigChangeContextOk returns a tuple with the ConfigChangeContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeContext

`func (o *ServerConfigChangeDetailRelationship) SetConfigChangeContext(v PolicyConfigResultContext)`

SetConfigChangeContext sets ConfigChangeContext field to given value.

### HasConfigChangeContext

`func (o *ServerConfigChangeDetailRelationship) HasConfigChangeContext() bool`

HasConfigChangeContext returns a boolean if a field has been set.

### GetConfigChangeFlag

`func (o *ServerConfigChangeDetailRelationship) GetConfigChangeFlag() string`

GetConfigChangeFlag returns the ConfigChangeFlag field if non-nil, zero value otherwise.

### GetConfigChangeFlagOk

`func (o *ServerConfigChangeDetailRelationship) GetConfigChangeFlagOk() (*string, bool)`

GetConfigChangeFlagOk returns a tuple with the ConfigChangeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigChangeFlag

`func (o *ServerConfigChangeDetailRelationship) SetConfigChangeFlag(v string)`

SetConfigChangeFlag sets ConfigChangeFlag field to given value.

### HasConfigChangeFlag

`func (o *ServerConfigChangeDetailRelationship) HasConfigChangeFlag() bool`

HasConfigChangeFlag returns a boolean if a field has been set.

### GetDisruptions

`func (o *ServerConfigChangeDetailRelationship) GetDisruptions() []string`

GetDisruptions returns the Disruptions field if non-nil, zero value otherwise.

### GetDisruptionsOk

`func (o *ServerConfigChangeDetailRelationship) GetDisruptionsOk() (*[]string, bool)`

GetDisruptionsOk returns a tuple with the Disruptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisruptions

`func (o *ServerConfigChangeDetailRelationship) SetDisruptions(v []string)`

SetDisruptions sets Disruptions field to given value.

### HasDisruptions

`func (o *ServerConfigChangeDetailRelationship) HasDisruptions() bool`

HasDisruptions returns a boolean if a field has been set.

### GetMessage

`func (o *ServerConfigChangeDetailRelationship) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServerConfigChangeDetailRelationship) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServerConfigChangeDetailRelationship) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServerConfigChangeDetailRelationship) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModStatus

`func (o *ServerConfigChangeDetailRelationship) GetModStatus() string`

GetModStatus returns the ModStatus field if non-nil, zero value otherwise.

### GetModStatusOk

`func (o *ServerConfigChangeDetailRelationship) GetModStatusOk() (*string, bool)`

GetModStatusOk returns a tuple with the ModStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModStatus

`func (o *ServerConfigChangeDetailRelationship) SetModStatus(v string)`

SetModStatus sets ModStatus field to given value.

### HasModStatus

`func (o *ServerConfigChangeDetailRelationship) HasModStatus() bool`

HasModStatus returns a boolean if a field has been set.

### GetProfile

`func (o *ServerConfigChangeDetailRelationship) GetProfile() ServerProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerConfigChangeDetailRelationship) GetProfileOk() (*ServerProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerConfigChangeDetailRelationship) SetProfile(v ServerProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerConfigChangeDetailRelationship) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


