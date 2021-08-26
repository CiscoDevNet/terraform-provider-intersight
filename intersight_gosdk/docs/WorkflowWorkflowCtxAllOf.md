# WorkflowWorkflowCtxAllOf

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

### NewWorkflowWorkflowCtxAllOf

`func NewWorkflowWorkflowCtxAllOf(classId string, objectType string, ) *WorkflowWorkflowCtxAllOf`

NewWorkflowWorkflowCtxAllOf instantiates a new WorkflowWorkflowCtxAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWorkflowCtxAllOfWithDefaults

`func NewWorkflowWorkflowCtxAllOfWithDefaults() *WorkflowWorkflowCtxAllOf`

NewWorkflowWorkflowCtxAllOfWithDefaults instantiates a new WorkflowWorkflowCtxAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowWorkflowCtxAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowWorkflowCtxAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowWorkflowCtxAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowWorkflowCtxAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowWorkflowCtxAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowWorkflowCtxAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


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

### SetInitiatorCtxNil

`func (o *WorkflowWorkflowCtxAllOf) SetInitiatorCtxNil(b bool)`

 SetInitiatorCtxNil sets the value for InitiatorCtx to be an explicit nil

### UnsetInitiatorCtx
`func (o *WorkflowWorkflowCtxAllOf) UnsetInitiatorCtx()`

UnsetInitiatorCtx ensures that no value is present for InitiatorCtx, not even an explicit nil
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

### SetTargetCtxListNil

`func (o *WorkflowWorkflowCtxAllOf) SetTargetCtxListNil(b bool)`

 SetTargetCtxListNil sets the value for TargetCtxList to be an explicit nil

### UnsetTargetCtxList
`func (o *WorkflowWorkflowCtxAllOf) UnsetTargetCtxList()`

UnsetTargetCtxList ensures that no value is present for TargetCtxList, not even an explicit nil
### GetWorkflowMetaName

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowMetaName() string`

GetWorkflowMetaName returns the WorkflowMetaName field if non-nil, zero value otherwise.

### GetWorkflowMetaNameOk

`func (o *WorkflowWorkflowCtxAllOf) GetWorkflowMetaNameOk() (*string, bool)`

GetWorkflowMetaNameOk returns a tuple with the WorkflowMetaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowMetaName

`func (o *WorkflowWorkflowCtxAllOf) SetWorkflowMetaName(v string)`

SetWorkflowMetaName sets WorkflowMetaName field to given value.

### HasWorkflowMetaName

`func (o *WorkflowWorkflowCtxAllOf) HasWorkflowMetaName() bool`

HasWorkflowMetaName returns a boolean if a field has been set.

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


