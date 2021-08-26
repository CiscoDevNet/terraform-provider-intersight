# WorkflowWorkflowCtx

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.WorkflowCtx"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.WorkflowCtx"]
**InitiatorCtx** | Pointer to [**NullableWorkflowInitiatorContext**](WorkflowInitiatorContext.md) |  | [optional] 
**TargetCtxList** | Pointer to [**[]WorkflowTargetContext**](WorkflowTargetContext.md) |  | [optional] 
**WorkflowMetaName** | Pointer to **string** | The name of workflowMeta of the workflow running. | [optional] 
**WorkflowSubtype** | Pointer to **string** | The subtype of the workflow. | [optional] 
**WorkflowType** | Pointer to **string** | Type of the workflow being started. This can be any string for client services to distinguish workflow by type. | [optional] 

## Methods

### NewWorkflowWorkflowCtx

`func NewWorkflowWorkflowCtx(classId string, objectType string, ) *WorkflowWorkflowCtx`

NewWorkflowWorkflowCtx instantiates a new WorkflowWorkflowCtx object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowCtxWithDefaults

`func NewWorkflowWorkflowCtxWithDefaults() *WorkflowWorkflowCtx`

NewWorkflowWorkflowCtxWithDefaults instantiates a new WorkflowWorkflowCtx object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowCtx) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowCtx) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowCtx) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowCtx) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowCtx) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowCtx) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetInitiatorCtxNil

`func (o *WorkflowWorkflowCtx) SetInitiatorCtxNil(b bool)`

 SetInitiatorCtxNil sets the value for InitiatorCtx to be an explicit nil

### UnsetInitiatorCtx
`func (o *WorkflowWorkflowCtx) UnsetInitiatorCtx()`

UnsetInitiatorCtx ensures that no value is present for InitiatorCtx, not even an explicit nil
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

### SetTargetCtxListNil

`func (o *WorkflowWorkflowCtx) SetTargetCtxListNil(b bool)`

 SetTargetCtxListNil sets the value for TargetCtxList to be an explicit nil

### UnsetTargetCtxList
`func (o *WorkflowWorkflowCtx) UnsetTargetCtxList()`

UnsetTargetCtxList ensures that no value is present for TargetCtxList, not even an explicit nil
### GetWorkflowMetaName

`func (o *WorkflowWorkflowCtx) GetWorkflowMetaName() string`

GetWorkflowMetaName returns the WorkflowMetaName field if non-nil, zero value otherwise.

### GetWorkflowMetaNameOk

`func (o *WorkflowWorkflowCtx) GetWorkflowMetaNameOk() (*string, bool)`

GetWorkflowMetaNameOk returns a tuple with the WorkflowMetaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetaName

`func (o *WorkflowWorkflowCtx) SetWorkflowMetaName(v string)`

SetWorkflowMetaName sets WorkflowMetaName field to given value.

### HasWorkflowMetaName

`func (o *WorkflowWorkflowCtx) HasWorkflowMetaName() bool`

HasWorkflowMetaName returns a boolean if a field has been set.

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


