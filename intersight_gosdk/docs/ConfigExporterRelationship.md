# ConfigExporterRelationship

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
**DownloadPath** | Pointer to **string** | Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired. | [optional] [readonly] 
**Items** | Pointer to [**[]ConfigMoRef**](config.MoRef.md) |  | [optional] 
**Name** | Pointer to **string** | An identifier for the exporter instance. | [optional] 
**Status** | Pointer to **string** | Status of the export operation. * &#x60;&#x60; - The operation has not started. * &#x60;InProgress&#x60; - The operation is in progress. * &#x60;Success&#x60; - The operation has succeeded. * &#x60;Failed&#x60; - The operation has failed. * &#x60;RollBackInitiated&#x60; - The rollback has been inititiated for import failure. * &#x60;RollBackFailed&#x60; - The rollback has failed for import failure. * &#x60;RollbackCompleted&#x60; - The rollback has completed for import failure. * &#x60;RollbackAborted&#x60; - The rollback has been aborted for import failure. * &#x60;OperationTimedOut&#x60; - The operation has timed out. | [optional] [readonly] [default to ""]
**StatusMessage** | Pointer to **string** | Status message associated with failures or progress indication. | [optional] [readonly] 
**ExportedItems** | Pointer to [**[]ConfigExportedItemRelationship**](config.ExportedItem.Relationship.md) | An array of relationships to configExportedItem resources. | [optional] [readonly] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 

## Methods

### NewConfigExporterRelationship

`func NewConfigExporterRelationship(classId string, objectType string, ) *ConfigExporterRelationship`

NewConfigExporterRelationship instantiates a new ConfigExporterRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExporterRelationshipWithDefaults

`func NewConfigExporterRelationshipWithDefaults() *ConfigExporterRelationship`

NewConfigExporterRelationshipWithDefaults instantiates a new ConfigExporterRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *ConfigExporterRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *ConfigExporterRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *ConfigExporterRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *ConfigExporterRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *ConfigExporterRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *ConfigExporterRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *ConfigExporterRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *ConfigExporterRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ConfigExporterRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ConfigExporterRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ConfigExporterRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *ConfigExporterRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *ConfigExporterRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *ConfigExporterRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *ConfigExporterRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *ConfigExporterRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *ConfigExporterRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *ConfigExporterRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *ConfigExporterRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *ConfigExporterRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *ConfigExporterRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *ConfigExporterRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *ConfigExporterRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *ConfigExporterRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ConfigExporterRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ConfigExporterRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *ConfigExporterRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *ConfigExporterRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *ConfigExporterRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *ConfigExporterRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *ConfigExporterRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *ConfigExporterRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *ConfigExporterRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *ConfigExporterRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *ConfigExporterRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigExporterRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigExporterRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigExporterRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *ConfigExporterRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *ConfigExporterRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *ConfigExporterRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *ConfigExporterRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *ConfigExporterRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *ConfigExporterRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *ConfigExporterRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *ConfigExporterRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *ConfigExporterRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *ConfigExporterRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *ConfigExporterRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ConfigExporterRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ConfigExporterRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ConfigExporterRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *ConfigExporterRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *ConfigExporterRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *ConfigExporterRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *ConfigExporterRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *ConfigExporterRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *ConfigExporterRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *ConfigExporterRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *ConfigExporterRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *ConfigExporterRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *ConfigExporterRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *ConfigExporterRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *ConfigExporterRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetDownloadPath

`func (o *ConfigExporterRelationship) GetDownloadPath() string`

GetDownloadPath returns the DownloadPath field if non-nil, zero value otherwise.

### GetDownloadPathOk

`func (o *ConfigExporterRelationship) GetDownloadPathOk() (*string, bool)`

GetDownloadPathOk returns a tuple with the DownloadPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPath

`func (o *ConfigExporterRelationship) SetDownloadPath(v string)`

SetDownloadPath sets DownloadPath field to given value.

### HasDownloadPath

`func (o *ConfigExporterRelationship) HasDownloadPath() bool`

HasDownloadPath returns a boolean if a field has been set.

### GetItems

`func (o *ConfigExporterRelationship) GetItems() []ConfigMoRef`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConfigExporterRelationship) GetItemsOk() (*[]ConfigMoRef, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConfigExporterRelationship) SetItems(v []ConfigMoRef)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConfigExporterRelationship) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *ConfigExporterRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigExporterRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigExporterRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigExporterRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ConfigExporterRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConfigExporterRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConfigExporterRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConfigExporterRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ConfigExporterRelationship) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ConfigExporterRelationship) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ConfigExporterRelationship) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ConfigExporterRelationship) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetExportedItems

`func (o *ConfigExporterRelationship) GetExportedItems() []ConfigExportedItemRelationship`

GetExportedItems returns the ExportedItems field if non-nil, zero value otherwise.

### GetExportedItemsOk

`func (o *ConfigExporterRelationship) GetExportedItemsOk() (*[]ConfigExportedItemRelationship, bool)`

GetExportedItemsOk returns a tuple with the ExportedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedItems

`func (o *ConfigExporterRelationship) SetExportedItems(v []ConfigExportedItemRelationship)`

SetExportedItems sets ExportedItems field to given value.

### HasExportedItems

`func (o *ConfigExporterRelationship) HasExportedItems() bool`

HasExportedItems returns a boolean if a field has been set.

### SetExportedItemsNil

`func (o *ConfigExporterRelationship) SetExportedItemsNil(b bool)`

 SetExportedItemsNil sets the value for ExportedItems to be an explicit nil

### UnsetExportedItems
`func (o *ConfigExporterRelationship) UnsetExportedItems()`

UnsetExportedItems ensures that no value is present for ExportedItems, not even an explicit nil
### GetOrganization

`func (o *ConfigExporterRelationship) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ConfigExporterRelationship) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ConfigExporterRelationship) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ConfigExporterRelationship) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


