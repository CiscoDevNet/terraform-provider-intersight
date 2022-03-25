# WorkflowTaskLoopInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TaskLoopInfo"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TaskLoopInfo"]
**Iteration** | Pointer to **int64** | This specifies the count of iteration for the specific task executed inside the loop. | [optional] 
**LoopTaskLabel** | Pointer to **string** | Label of the loop task inside which this task is executed. | [optional] 
**LoopTaskName** | Pointer to **string** | Name of the loop task inside which this task is executed. | [optional] 
**LoopType** | Pointer to **string** | This specifies the type of loop, Serial or Parallel. * &#x60;None&#x60; - The enum specifies the option as None which implies this is not a Loop type and this is the default value for loop type. * &#x60;Parallel&#x60; - The enum specifies the option as Parallel where the loop task type is parallel loop. * &#x60;Serial&#x60; - The enum specifies the option as Serial where the loop task type is serial loop. | [optional] [default to "None"]

## Methods

### NewWorkflowTaskLoopInfoAllOf

`func NewWorkflowTaskLoopInfoAllOf(classId string, objectType string, ) *WorkflowTaskLoopInfoAllOf`

NewWorkflowTaskLoopInfoAllOf instantiates a new WorkflowTaskLoopInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTaskLoopInfoAllOfWithDefaults

`func NewWorkflowTaskLoopInfoAllOfWithDefaults() *WorkflowTaskLoopInfoAllOf`

NewWorkflowTaskLoopInfoAllOfWithDefaults instantiates a new WorkflowTaskLoopInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTaskLoopInfoAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTaskLoopInfoAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTaskLoopInfoAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTaskLoopInfoAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTaskLoopInfoAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTaskLoopInfoAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetIteration

`func (o *WorkflowTaskLoopInfoAllOf) GetIteration() int64`

GetIteration returns the Iteration field if non-nil, zero value otherwise.

### GetIterationOk

`func (o *WorkflowTaskLoopInfoAllOf) GetIterationOk() (*int64, bool)`

GetIterationOk returns a tuple with the Iteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIteration

`func (o *WorkflowTaskLoopInfoAllOf) SetIteration(v int64)`

SetIteration sets Iteration field to given value.

### HasIteration

`func (o *WorkflowTaskLoopInfoAllOf) HasIteration() bool`

HasIteration returns a boolean if a field has been set.

### GetLoopTaskLabel

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskLabel() string`

GetLoopTaskLabel returns the LoopTaskLabel field if non-nil, zero value otherwise.

### GetLoopTaskLabelOk

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskLabelOk() (*string, bool)`

GetLoopTaskLabelOk returns a tuple with the LoopTaskLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopTaskLabel

`func (o *WorkflowTaskLoopInfoAllOf) SetLoopTaskLabel(v string)`

SetLoopTaskLabel sets LoopTaskLabel field to given value.

### HasLoopTaskLabel

`func (o *WorkflowTaskLoopInfoAllOf) HasLoopTaskLabel() bool`

HasLoopTaskLabel returns a boolean if a field has been set.

### GetLoopTaskName

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskName() string`

GetLoopTaskName returns the LoopTaskName field if non-nil, zero value otherwise.

### GetLoopTaskNameOk

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopTaskNameOk() (*string, bool)`

GetLoopTaskNameOk returns a tuple with the LoopTaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopTaskName

`func (o *WorkflowTaskLoopInfoAllOf) SetLoopTaskName(v string)`

SetLoopTaskName sets LoopTaskName field to given value.

### HasLoopTaskName

`func (o *WorkflowTaskLoopInfoAllOf) HasLoopTaskName() bool`

HasLoopTaskName returns a boolean if a field has been set.

### GetLoopType

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopType() string`

GetLoopType returns the LoopType field if non-nil, zero value otherwise.

### GetLoopTypeOk

`func (o *WorkflowTaskLoopInfoAllOf) GetLoopTypeOk() (*string, bool)`

GetLoopTypeOk returns a tuple with the LoopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopType

`func (o *WorkflowTaskLoopInfoAllOf) SetLoopType(v string)`

SetLoopType sets LoopType field to given value.

### HasLoopType

`func (o *WorkflowTaskLoopInfoAllOf) HasLoopType() bool`

HasLoopType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


