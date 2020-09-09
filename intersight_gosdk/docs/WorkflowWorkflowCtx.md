# WorkflowWorkflowCtx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorCtx** | Pointer to [**WorkflowInitiatorContext**](workflow.InitiatorContext.md) |  | [optional] 
**TargetCtxList** | Pointer to [**[]WorkflowTargetContext**](workflow.TargetContext.md) |  | [optional] 
**WorkflowSubtype** | Pointer to **string** | The subtype of the workflow. | [optional] 
**WorkflowType** | Pointer to **string** | Type of the workflow being started. This can be any string for client services to distinguish workflow by type. | [optional] 

## Methods

### NewWorkflowWorkflowCtx

`func NewWorkflowWorkflowCtx() *WorkflowWorkflowCtx`

NewWorkflowWorkflowCtx instantiates a new WorkflowWorkflowCtx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowCtxWithDefaults

`func NewWorkflowWorkflowCtxWithDefaults() *WorkflowWorkflowCtx`

NewWorkflowWorkflowCtxWithDefaults instantiates a new WorkflowWorkflowCtx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorCtx

`func (o *WorkflowWorkflowCtx) GetInitiatorCtx() WorkflowInitiatorContext`

GetInitiatorCtx returns the InitiatorCtx field if non-nil, zero value otherwise.

### GetInitiatorCtxOk

`func (o *WorkflowWorkflowCtx) GetInitiatorCtxOk() (*WorkflowInitiatorContext, bool)`

GetInitiatorCtxOk returns a tuple with the InitiatorCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCtx

`func (o *WorkflowWorkflowCtx) SetInitiatorCtx(v WorkflowInitiatorContext)`

SetInitiatorCtx sets InitiatorCtx field to given value.

### HasInitiatorCtx

`func (o *WorkflowWorkflowCtx) HasInitiatorCtx() bool`

HasInitiatorCtx returns a boolean if a field has been set.

### GetTargetCtxList

`func (o *WorkflowWorkflowCtx) GetTargetCtxList() []WorkflowTargetContext`

GetTargetCtxList returns the TargetCtxList field if non-nil, zero value otherwise.

### GetTargetCtxListOk

`func (o *WorkflowWorkflowCtx) GetTargetCtxListOk() (*[]WorkflowTargetContext, bool)`

GetTargetCtxListOk returns a tuple with the TargetCtxList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCtxList

`func (o *WorkflowWorkflowCtx) SetTargetCtxList(v []WorkflowTargetContext)`

SetTargetCtxList sets TargetCtxList field to given value.

### HasTargetCtxList

`func (o *WorkflowWorkflowCtx) HasTargetCtxList() bool`

HasTargetCtxList returns a boolean if a field has been set.

### GetWorkflowSubtype

`func (o *WorkflowWorkflowCtx) GetWorkflowSubtype() string`

GetWorkflowSubtype returns the WorkflowSubtype field if non-nil, zero value otherwise.

### GetWorkflowSubtypeOk

`func (o *WorkflowWorkflowCtx) GetWorkflowSubtypeOk() (*string, bool)`

GetWorkflowSubtypeOk returns a tuple with the WorkflowSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSubtype

`func (o *WorkflowWorkflowCtx) SetWorkflowSubtype(v string)`

SetWorkflowSubtype sets WorkflowSubtype field to given value.

### HasWorkflowSubtype

`func (o *WorkflowWorkflowCtx) HasWorkflowSubtype() bool`

HasWorkflowSubtype returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowWorkflowCtx) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowWorkflowCtx) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowWorkflowCtx) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowWorkflowCtx) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


