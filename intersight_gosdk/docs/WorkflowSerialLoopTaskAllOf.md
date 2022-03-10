# WorkflowSerialLoopTaskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.SerialLoopTask"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.SerialLoopTask"]
**Condition** | Pointer to **string** | Condition expression which will be evaluated and when expression is true the tasks under loop will be executed. | [optional] 

## Methods

### NewWorkflowSerialLoopTaskAllOf

`func NewWorkflowSerialLoopTaskAllOf(classId string, objectType string, ) *WorkflowSerialLoopTaskAllOf`

NewWorkflowSerialLoopTaskAllOf instantiates a new WorkflowSerialLoopTaskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSerialLoopTaskAllOfWithDefaults

`func NewWorkflowSerialLoopTaskAllOfWithDefaults() *WorkflowSerialLoopTaskAllOf`

NewWorkflowSerialLoopTaskAllOfWithDefaults instantiates a new WorkflowSerialLoopTaskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowSerialLoopTaskAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowSerialLoopTaskAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowSerialLoopTaskAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowSerialLoopTaskAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowSerialLoopTaskAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowSerialLoopTaskAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetCondition

`func (o *WorkflowSerialLoopTaskAllOf) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *WorkflowSerialLoopTaskAllOf) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *WorkflowSerialLoopTaskAllOf) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *WorkflowSerialLoopTaskAllOf) HasCondition() bool`

HasCondition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


