# WorkflowWorkflowCtxAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitiatorCtx** | Pointer to [**WorkflowInitiatorContext**](workflow.InitiatorContext.md) |  | [optional] 
**TargetCtxList** | Pointer to [**[]WorkflowTargetContext**](workflow.TargetContext.md) |  | [optional] 
**WorkflowSubtype** | Pointer to **string** | The subtype of the workflow. | [optional] 
**WorkflowType** | Pointer to **string** | Type of the workflow being started. This can be any string for client services to distinguish workflow by type. | [optional] 

## Methods

### NewWorkflowWorkflowCtxAllOf

`func NewWorkflowWorkflowCtxAllOf() *WorkflowWorkflowCtxAllOf`

NewWorkflowWorkflowCtxAllOf instantiates a new WorkflowWorkflowCtxAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowCtxAllOfWithDefaults

`func NewWorkflowWorkflowCtxAllOfWithDefaults() *WorkflowWorkflowCtxAllOf`

NewWorkflowWorkflowCtxAllOfWithDefaults instantiates a new WorkflowWorkflowCtxAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitiatorCtx

`func (o *WorkflowWorkflowCtxAllOf) GetInitiatorCtx() WorkflowInitiatorContext`

GetInitiatorCtx returns the InitiatorCtx field if non-nil, zero value otherwise.

### GetInitiatorCtxOk

`func (o *WorkflowWorkflowCtxAllOf) GetInitiatorCtxOk() (*WorkflowInitiatorContext, bool)`

GetInitiatorCtxOk returns a tuple with the InitiatorCtx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCtx

`func (o *WorkflowWorkflowCtxAllOf) SetInitiatorCtx(v WorkflowInitiatorContext)`

SetInitiatorCtx sets InitiatorCtx field to given value.

### HasInitiatorCtx

`func (o *WorkflowWorkflowCtxAllOf) HasInitiatorCtx() bool`

HasInitiatorCtx returns a boolean if a field has been set.

### GetTargetCtxList

`func (o *WorkflowWorkflowCtxAllOf) GetTargetCtxList() []WorkflowTargetContext`

GetTargetCtxList returns the TargetCtxList field if non-nil, zero value otherwise.

### GetTargetCtxListOk

`func (o *WorkflowWorkflowCtxAllOf) GetTargetCtxListOk() (*[]WorkflowTargetContext, bool)`

GetTargetCtxListOk returns a tuple with the TargetCtxList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCtxList

`func (o *WorkflowWorkflowCtxAllOf) SetTargetCtxList(v []WorkflowTargetContext)`

SetTargetCtxList sets TargetCtxList field to given value.

### HasTargetCtxList

`func (o *WorkflowWorkflowCtxAllOf) HasTargetCtxList() bool`

HasTargetCtxList returns a boolean if a field has been set.

### GetWorkflowSubtype

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowSubtype() string`

GetWorkflowSubtype returns the WorkflowSubtype field if non-nil, zero value otherwise.

### GetWorkflowSubtypeOk

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowSubtypeOk() (*string, bool)`

GetWorkflowSubtypeOk returns a tuple with the WorkflowSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSubtype

`func (o *WorkflowWorkflowCtxAllOf) SetWorkflowSubtype(v string)`

SetWorkflowSubtype sets WorkflowSubtype field to given value.

### HasWorkflowSubtype

`func (o *WorkflowWorkflowCtxAllOf) HasWorkflowSubtype() bool`

HasWorkflowSubtype returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowWorkflowCtxAllOf) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowWorkflowCtxAllOf) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


