# WorkflowPendingDynamicWorkflowInfoRelationship

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
**Input** | Pointer to **interface{}** | The inputs of the workflow. | [optional] 
**Name** | Pointer to **string** | A name for the pending dynamic workflow. | [optional] 
**PendingServices** | Pointer to **[]string** |  | [optional] 
**Src** | Pointer to **string** | The src is workflow owner service. | [optional] 
**Status** | Pointer to **string** | The current status of the PendingDynamicWorkflowInfo. * &#x60;GatheringTasks&#x60; - Dynamic workflow is gathering tasks before workflow can start execution. * &#x60;Waiting&#x60; - Dynamic workflow is in waiting state and not yet started execution. | [optional] [default to "GatheringTasks"]
**WaitOnDuplicate** | Pointer to **bool** | When set to true workflow engine will wait for a duplicate to finish before starting a new one. | [optional] 
**WorkflowActionTaskLists** | Pointer to [**[]WorkflowDynamicWorkflowActionTaskList**](workflow.DynamicWorkflowActionTaskList.md) |  | [optional] 
**WorkflowCtx** | Pointer to [**WorkflowWorkflowCtx**](workflow.WorkflowCtx.md) |  | [optional] 
**WorkflowKey** | Pointer to **string** | This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup. | [optional] 
**WorkflowMeta** | Pointer to **interface{}** | The metadata of the workflow. | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowPendingDynamicWorkflowInfoRelationship

`func NewWorkflowPendingDynamicWorkflowInfoRelationship(classId string, objectType string, ) *WorkflowPendingDynamicWorkflowInfoRelationship`

NewWorkflowPendingDynamicWorkflowInfoRelationship instantiates a new WorkflowPendingDynamicWorkflowInfoRelationship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPendingDynamicWorkflowInfoRelationshipWithDefaults

`func NewWorkflowPendingDynamicWorkflowInfoRelationshipWithDefaults() *WorkflowPendingDynamicWorkflowInfoRelationship`

NewWorkflowPendingDynamicWorkflowInfoRelationshipWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfoRelationship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetAccountMoid() string`

GetAccountMoid returns the AccountMoid field if non-nil, zero value otherwise.

### GetAccountMoidOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetAccountMoidOk() (*string, bool)`

GetAccountMoidOk returns a tuple with the AccountMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetAccountMoid(v string)`

SetAccountMoid sets AccountMoid field to given value.

### HasAccountMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasAccountMoid() bool`

HasAccountMoid returns a boolean if a field has been set.

### GetClassId

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetCreateTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDomainGroupMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetDomainGroupMoid() string`

GetDomainGroupMoid returns the DomainGroupMoid field if non-nil, zero value otherwise.

### GetDomainGroupMoidOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetDomainGroupMoidOk() (*string, bool)`

GetDomainGroupMoidOk returns a tuple with the DomainGroupMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainGroupMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetDomainGroupMoid(v string)`

SetDomainGroupMoid sets DomainGroupMoid field to given value.

### HasDomainGroupMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasDomainGroupMoid() bool`

HasDomainGroupMoid returns a boolean if a field has been set.

### GetModTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetModTime() time.Time`

GetModTime returns the ModTime field if non-nil, zero value otherwise.

### GetModTimeOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetModTimeOk() (*time.Time, bool)`

GetModTimeOk returns a tuple with the ModTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetModTime(v time.Time)`

SetModTime sets ModTime field to given value.

### HasModTime

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasModTime() bool`

HasModTime returns a boolean if a field has been set.

### GetMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetMoid() string`

GetMoid returns the Moid field if non-nil, zero value otherwise.

### GetMoidOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetMoidOk() (*string, bool)`

GetMoidOk returns a tuple with the Moid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetMoid(v string)`

SetMoid sets Moid field to given value.

### HasMoid

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasMoid() bool`

HasMoid returns a boolean if a field has been set.

### GetObjectType

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetOwners

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetOwners() []string`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetOwnersOk() (*[]string, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetOwners(v []string)`

SetOwners sets Owners field to given value.

### HasOwners

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasOwners() bool`

HasOwners returns a boolean if a field has been set.

### GetSharedScope

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetSharedScope() string`

GetSharedScope returns the SharedScope field if non-nil, zero value otherwise.

### GetSharedScopeOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetSharedScopeOk() (*string, bool)`

GetSharedScopeOk returns a tuple with the SharedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedScope

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetSharedScope(v string)`

SetSharedScope sets SharedScope field to given value.

### HasSharedScope

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasSharedScope() bool`

HasSharedScope returns a boolean if a field has been set.

### GetTags

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetTags() []MoTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetTagsOk() (*[]MoTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetTags(v []MoTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersionContext

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetVersionContext() MoVersionContext`

GetVersionContext returns the VersionContext field if non-nil, zero value otherwise.

### GetVersionContextOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetVersionContextOk() (*MoVersionContext, bool)`

GetVersionContextOk returns a tuple with the VersionContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionContext

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetVersionContext(v MoVersionContext)`

SetVersionContext sets VersionContext field to given value.

### HasVersionContext

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasVersionContext() bool`

HasVersionContext returns a boolean if a field has been set.

### GetAncestors

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetAncestors() []MoBaseMoRelationship`

GetAncestors returns the Ancestors field if non-nil, zero value otherwise.

### GetAncestorsOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetAncestorsOk() (*[]MoBaseMoRelationship, bool)`

GetAncestorsOk returns a tuple with the Ancestors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAncestors

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetAncestors(v []MoBaseMoRelationship)`

SetAncestors sets Ancestors field to given value.

### HasAncestors

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasAncestors() bool`

HasAncestors returns a boolean if a field has been set.

### SetAncestorsNil

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetAncestorsNil(b bool)`

 SetAncestorsNil sets the value for Ancestors to be an explicit nil

### UnsetAncestors
`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) UnsetAncestors()`

UnsetAncestors ensures that no value is present for Ancestors, not even an explicit nil
### GetParent

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetParent() MoBaseMoRelationship`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetParentOk() (*MoBaseMoRelationship, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetParent(v MoBaseMoRelationship)`

SetParent sets Parent field to given value.

### HasParent

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPermissionResources

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetPermissionResources() []MoBaseMoRelationship`

GetPermissionResources returns the PermissionResources field if non-nil, zero value otherwise.

### GetPermissionResourcesOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetPermissionResourcesOk() (*[]MoBaseMoRelationship, bool)`

GetPermissionResourcesOk returns a tuple with the PermissionResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionResources

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetPermissionResources(v []MoBaseMoRelationship)`

SetPermissionResources sets PermissionResources field to given value.

### HasPermissionResources

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasPermissionResources() bool`

HasPermissionResources returns a boolean if a field has been set.

### SetPermissionResourcesNil

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetPermissionResourcesNil(b bool)`

 SetPermissionResourcesNil sets the value for PermissionResources to be an explicit nil

### UnsetPermissionResources
`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) UnsetPermissionResources()`

UnsetPermissionResources ensures that no value is present for PermissionResources, not even an explicit nil
### GetDisplayNames

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetDisplayNames() map[string][]string`

GetDisplayNames returns the DisplayNames field if non-nil, zero value otherwise.

### GetDisplayNamesOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetDisplayNamesOk() (*map[string][]string, bool)`

GetDisplayNamesOk returns a tuple with the DisplayNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayNames

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetDisplayNames(v map[string][]string)`

SetDisplayNames sets DisplayNames field to given value.

### HasDisplayNames

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasDisplayNames() bool`

HasDisplayNames returns a boolean if a field has been set.

### SetDisplayNamesNil

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetDisplayNamesNil(b bool)`

 SetDisplayNamesNil sets the value for DisplayNames to be an explicit nil

### UnsetDisplayNames
`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) UnsetDisplayNames()`

UnsetDisplayNames ensures that no value is present for DisplayNames, not even an explicit nil
### GetInput

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetPendingServices() []string`

GetPendingServices returns the PendingServices field if non-nil, zero value otherwise.

### GetPendingServicesOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetPendingServicesOk() (*[]string, bool)`

GetPendingServicesOk returns a tuple with the PendingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetPendingServices(v []string)`

SetPendingServices sets PendingServices field to given value.

### HasPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasPendingServices() bool`

HasPendingServices returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList`

GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field if non-nil, zero value otherwise.

### GetWorkflowActionTaskListsOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowActionTaskListsOk() (*[]WorkflowDynamicWorkflowActionTaskList, bool)`

GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList)`

SetWorkflowActionTaskLists sets WorkflowActionTaskLists field to given value.

### HasWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWorkflowActionTaskLists() bool`

HasWorkflowActionTaskLists returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### GetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowKey() string`

GetWorkflowKey returns the WorkflowKey field if non-nil, zero value otherwise.

### GetWorkflowKeyOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowKeyOk() (*string, bool)`

GetWorkflowKeyOk returns a tuple with the WorkflowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowKey(v string)`

SetWorkflowKey sets WorkflowKey field to given value.

### HasWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWorkflowKey() bool`

HasWorkflowKey returns a boolean if a field has been set.

### GetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowMeta() interface{}`

GetWorkflowMeta returns the WorkflowMeta field if non-nil, zero value otherwise.

### GetWorkflowMetaOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowMetaOk() (*interface{}, bool)`

GetWorkflowMetaOk returns a tuple with the WorkflowMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowMeta(v interface{})`

SetWorkflowMeta sets WorkflowMeta field to given value.

### HasWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWorkflowMeta() bool`

HasWorkflowMeta returns a boolean if a field has been set.

### SetWorkflowMetaNil

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowMetaNil(b bool)`

 SetWorkflowMetaNil sets the value for WorkflowMeta to be an explicit nil

### UnsetWorkflowMeta
`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) UnsetWorkflowMeta()`

UnsetWorkflowMeta ensures that no value is present for WorkflowMeta, not even an explicit nil
### GetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoRelationship) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


