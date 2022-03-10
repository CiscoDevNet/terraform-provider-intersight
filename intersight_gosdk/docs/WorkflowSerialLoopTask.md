# WorkflowSerialLoopTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SerialLoopTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SerialLoopTask"]
**Condition** | Pointer to **string** | Condition expression which will be evaluated and when expression is true the tasks under loop will be executed. | [optional] 

## Methods

### NewWorkflowSerialLoopTask

`func NewWorkflowSerialLoopTask(classId string, objectType string, ) *WorkflowSerialLoopTask`

NewWorkflowSerialLoopTask instantiates a new WorkflowSerialLoopTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSerialLoopTaskWithDefaults

`func NewWorkflowSerialLoopTaskWithDefaults() *WorkflowSerialLoopTask`

NewWorkflowSerialLoopTaskWithDefaults instantiates a new WorkflowSerialLoopTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSerialLoopTask) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSerialLoopTask) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSerialLoopTask) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSerialLoopTask) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSerialLoopTask) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSerialLoopTask) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *WorkflowSerialLoopTask) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WorkflowSerialLoopTask) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WorkflowSerialLoopTask) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WorkflowSerialLoopTask) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


