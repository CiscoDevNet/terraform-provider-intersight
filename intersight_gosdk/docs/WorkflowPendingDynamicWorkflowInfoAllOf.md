# WorkflowPendingDynamicWorkflowInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewWorkflowPendingDynamicWorkflowInfoAllOf

`func NewWorkflowPendingDynamicWorkflowInfoAllOf() *WorkflowPendingDynamicWorkflowInfoAllOf`

NewWorkflowPendingDynamicWorkflowInfoAllOf instantiates a new WorkflowPendingDynamicWorkflowInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPendingDynamicWorkflowInfoAllOfWithDefaults

`func NewWorkflowPendingDynamicWorkflowInfoAllOfWithDefaults() *WorkflowPendingDynamicWorkflowInfoAllOf`

NewWorkflowPendingDynamicWorkflowInfoAllOfWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetPendingServices() []string`

GetPendingServices returns the PendingServices field if non-nil, zero value otherwise.

### GetPendingServicesOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetPendingServicesOk() (*[]string, bool)`

GetPendingServicesOk returns a tuple with the PendingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetPendingServices(v []string)`

SetPendingServices sets PendingServices field to given value.

### HasPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasPendingServices() bool`

HasPendingServices returns a boolean if a field has been set.

### GetSrc

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList`

GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field if non-nil, zero value otherwise.

### GetWorkflowActionTaskListsOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowActionTaskListsOk() (*[]WorkflowDynamicWorkflowActionTaskList, bool)`

GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList)`

SetWorkflowActionTaskLists sets WorkflowActionTaskLists field to given value.

### HasWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowActionTaskLists() bool`

HasWorkflowActionTaskLists returns a boolean if a field has been set.

### GetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### GetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowKey() string`

GetWorkflowKey returns the WorkflowKey field if non-nil, zero value otherwise.

### GetWorkflowKeyOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowKeyOk() (*string, bool)`

GetWorkflowKeyOk returns a tuple with the WorkflowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowKey(v string)`

SetWorkflowKey sets WorkflowKey field to given value.

### HasWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowKey() bool`

HasWorkflowKey returns a boolean if a field has been set.

### GetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowMeta() interface{}`

GetWorkflowMeta returns the WorkflowMeta field if non-nil, zero value otherwise.

### GetWorkflowMetaOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowMetaOk() (*interface{}, bool)`

GetWorkflowMetaOk returns a tuple with the WorkflowMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowMeta(v interface{})`

SetWorkflowMeta sets WorkflowMeta field to given value.

### HasWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowMeta() bool`

HasWorkflowMeta returns a boolean if a field has been set.

### SetWorkflowMetaNil

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowMetaNil(b bool)`

 SetWorkflowMetaNil sets the value for WorkflowMeta to be an explicit nil

### UnsetWorkflowMeta
`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) UnsetWorkflowMeta()`

UnsetWorkflowMeta ensures that no value is present for WorkflowMeta, not even an explicit nil
### GetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfoAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


