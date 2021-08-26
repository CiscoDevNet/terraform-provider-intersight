# WorkflowPendingDynamicWorkflowInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.PendingDynamicWorkflowInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.PendingDynamicWorkflowInfo"]
**Input** | Pointer to **interface{}** | The inputs of the workflow. | [optional] 
**Name** | Pointer to **string** | A name for the pending dynamic workflow. | [optional] 
**PendingServices** | Pointer to **[]string** |  | [optional] 
**Src** | Pointer to **string** | The src is workflow owner service. | [optional] 
**Status** | Pointer to **string** | The current status of the PendingDynamicWorkflowInfo. * &#x60;GatheringTasks&#x60; - Dynamic workflow is gathering tasks before workflow can start execution. * &#x60;Waiting&#x60; - Dynamic workflow is in waiting state and not yet started execution. | [optional] [default to "GatheringTasks"]
**WaitOnDuplicate** | Pointer to **bool** | When set to true workflow engine will wait for a duplicate to finish before starting a new one. | [optional] 
**WorkflowActionTaskLists** | Pointer to [**[]WorkflowDynamicWorkflowActionTaskList**](WorkflowDynamicWorkflowActionTaskList.md) |  | [optional] 
**WorkflowCtx** | Pointer to [**NullableWorkflowWorkflowCtx**](WorkflowWorkflowCtx.md) |  | [optional] 
**WorkflowKey** | Pointer to **string** | This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup. | [optional] 
**WorkflowMeta** | Pointer to **interface{}** | The metadata of the workflow. | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](WorkflowWorkflowInfoRelationship.md) |  | [optional] 

## Methods

### NewWorkflowPendingDynamicWorkflowInfo

`func NewWorkflowPendingDynamicWorkflowInfo(classId string, objectType string, ) *WorkflowPendingDynamicWorkflowInfo`

NewWorkflowPendingDynamicWorkflowInfo instantiates a new WorkflowPendingDynamicWorkflowInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPendingDynamicWorkflowInfoWithDefaults

`func NewWorkflowPendingDynamicWorkflowInfoWithDefaults() *WorkflowPendingDynamicWorkflowInfo`

NewWorkflowPendingDynamicWorkflowInfoWithDefaults instantiates a new WorkflowPendingDynamicWorkflowInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowPendingDynamicWorkflowInfo) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowPendingDynamicWorkflowInfo) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowPendingDynamicWorkflowInfo) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowPendingDynamicWorkflowInfo) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetInput

`func (o *WorkflowPendingDynamicWorkflowInfo) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *WorkflowPendingDynamicWorkflowInfo) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *WorkflowPendingDynamicWorkflowInfo) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *WorkflowPendingDynamicWorkflowInfo) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *WorkflowPendingDynamicWorkflowInfo) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetName

`func (o *WorkflowPendingDynamicWorkflowInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowPendingDynamicWorkflowInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowPendingDynamicWorkflowInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfo) GetPendingServices() []string`

GetPendingServices returns the PendingServices field if non-nil, zero value otherwise.

### GetPendingServicesOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetPendingServicesOk() (*[]string, bool)`

GetPendingServicesOk returns a tuple with the PendingServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfo) SetPendingServices(v []string)`

SetPendingServices sets PendingServices field to given value.

### HasPendingServices

`func (o *WorkflowPendingDynamicWorkflowInfo) HasPendingServices() bool`

HasPendingServices returns a boolean if a field has been set.

### SetPendingServicesNil

`func (o *WorkflowPendingDynamicWorkflowInfo) SetPendingServicesNil(b bool)`

 SetPendingServicesNil sets the value for PendingServices to be an explicit nil

### UnsetPendingServices
`func (o *WorkflowPendingDynamicWorkflowInfo) UnsetPendingServices()`

UnsetPendingServices ensures that no value is present for PendingServices, not even an explicit nil
### GetSrc

`func (o *WorkflowPendingDynamicWorkflowInfo) GetSrc() string`

GetSrc returns the Src field if non-nil, zero value otherwise.

### GetSrcOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetSrcOk() (*string, bool)`

GetSrcOk returns a tuple with the Src field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrc

`func (o *WorkflowPendingDynamicWorkflowInfo) SetSrc(v string)`

SetSrc sets Src field to given value.

### HasSrc

`func (o *WorkflowPendingDynamicWorkflowInfo) HasSrc() bool`

HasSrc returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowPendingDynamicWorkflowInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowPendingDynamicWorkflowInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowPendingDynamicWorkflowInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskLists() []WorkflowDynamicWorkflowActionTaskList`

GetWorkflowActionTaskLists returns the WorkflowActionTaskLists field if non-nil, zero value otherwise.

### GetWorkflowActionTaskListsOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowActionTaskListsOk() (*[]WorkflowDynamicWorkflowActionTaskList, bool)`

GetWorkflowActionTaskListsOk returns a tuple with the WorkflowActionTaskLists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowActionTaskLists(v []WorkflowDynamicWorkflowActionTaskList)`

SetWorkflowActionTaskLists sets WorkflowActionTaskLists field to given value.

### HasWorkflowActionTaskLists

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowActionTaskLists() bool`

HasWorkflowActionTaskLists returns a boolean if a field has been set.

### SetWorkflowActionTaskListsNil

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowActionTaskListsNil(b bool)`

 SetWorkflowActionTaskListsNil sets the value for WorkflowActionTaskLists to be an explicit nil

### UnsetWorkflowActionTaskLists
`func (o *WorkflowPendingDynamicWorkflowInfo) UnsetWorkflowActionTaskLists()`

UnsetWorkflowActionTaskLists ensures that no value is present for WorkflowActionTaskLists, not even an explicit nil
### GetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtx() WorkflowWorkflowCtx`

GetWorkflowCtx returns the WorkflowCtx field if non-nil, zero value otherwise.

### GetWorkflowCtxOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowCtxOk() (*WorkflowWorkflowCtx, bool)`

GetWorkflowCtxOk returns a tuple with the WorkflowCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowCtx(v WorkflowWorkflowCtx)`

SetWorkflowCtx sets WorkflowCtx field to given value.

### HasWorkflowCtx

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowCtx() bool`

HasWorkflowCtx returns a boolean if a field has been set.

### SetWorkflowCtxNil

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowCtxNil(b bool)`

 SetWorkflowCtxNil sets the value for WorkflowCtx to be an explicit nil

### UnsetWorkflowCtx
`func (o *WorkflowPendingDynamicWorkflowInfo) UnsetWorkflowCtx()`

UnsetWorkflowCtx ensures that no value is present for WorkflowCtx, not even an explicit nil
### GetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowKey() string`

GetWorkflowKey returns the WorkflowKey field if non-nil, zero value otherwise.

### GetWorkflowKeyOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowKeyOk() (*string, bool)`

GetWorkflowKeyOk returns a tuple with the WorkflowKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowKey(v string)`

SetWorkflowKey sets WorkflowKey field to given value.

### HasWorkflowKey

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowKey() bool`

HasWorkflowKey returns a boolean if a field has been set.

### GetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMeta() interface{}`

GetWorkflowMeta returns the WorkflowMeta field if non-nil, zero value otherwise.

### GetWorkflowMetaOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowMetaOk() (*interface{}, bool)`

GetWorkflowMetaOk returns a tuple with the WorkflowMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowMeta(v interface{})`

SetWorkflowMeta sets WorkflowMeta field to given value.

### HasWorkflowMeta

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowMeta() bool`

HasWorkflowMeta returns a boolean if a field has been set.

### SetWorkflowMetaNil

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowMetaNil(b bool)`

 SetWorkflowMetaNil sets the value for WorkflowMeta to be an explicit nil

### UnsetWorkflowMeta
`func (o *WorkflowPendingDynamicWorkflowInfo) UnsetWorkflowMeta()`

UnsetWorkflowMeta ensures that no value is present for WorkflowMeta, not even an explicit nil
### GetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *WorkflowPendingDynamicWorkflowInfo) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfo) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *WorkflowPendingDynamicWorkflowInfo) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


