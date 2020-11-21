# WorkflowValidationErrorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.ValidationError"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.ValidationError"]
**ErrorLog** | Pointer to **string** | Description of the error. | [optional] [readonly] 
**Field** | Pointer to **string** | When populated this refers to the input or output field within the workflow or task. | [optional] [readonly] 
**TaskName** | Pointer to **string** | The task name on which the error is found, when empty the error applies to the top level workflow. | [optional] [readonly] 
**TransitionName** | Pointer to **string** | When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue. | [optional] [readonly] 

## Methods

### NewWorkflowValidationErrorAllOf

`func NewWorkflowValidationErrorAllOf(classId string, objectType string, ) *WorkflowValidationErrorAllOf`

NewWorkflowValidationErrorAllOf instantiates a new WorkflowValidationErrorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowValidationErrorAllOfWithDefaults

`func NewWorkflowValidationErrorAllOfWithDefaults() *WorkflowValidationErrorAllOf`

NewWorkflowValidationErrorAllOfWithDefaults instantiates a new WorkflowValidationErrorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowValidationErrorAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowValidationErrorAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowValidationErrorAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowValidationErrorAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowValidationErrorAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowValidationErrorAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetErrorLog

`func (o *WorkflowValidationErrorAllOf) GetErrorLog() string`

GetErrorLog returns the ErrorLog field if non-nil, zero value otherwise.

### GetErrorLogOk

`func (o *WorkflowValidationErrorAllOf) GetErrorLogOk() (*string, bool)`

GetErrorLogOk returns a tuple with the ErrorLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorLog

`func (o *WorkflowValidationErrorAllOf) SetErrorLog(v string)`

SetErrorLog sets ErrorLog field to given value.

### HasErrorLog

`func (o *WorkflowValidationErrorAllOf) HasErrorLog() bool`

HasErrorLog returns a boolean if a field has been set.

### GetField

`func (o *WorkflowValidationErrorAllOf) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *WorkflowValidationErrorAllOf) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *WorkflowValidationErrorAllOf) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *WorkflowValidationErrorAllOf) HasField() bool`

HasField returns a boolean if a field has been set.

### GetTaskName

`func (o *WorkflowValidationErrorAllOf) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *WorkflowValidationErrorAllOf) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *WorkflowValidationErrorAllOf) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *WorkflowValidationErrorAllOf) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTransitionName

`func (o *WorkflowValidationErrorAllOf) GetTransitionName() string`

GetTransitionName returns the TransitionName field if non-nil, zero value otherwise.

### GetTransitionNameOk

`func (o *WorkflowValidationErrorAllOf) GetTransitionNameOk() (*string, bool)`

GetTransitionNameOk returns a tuple with the TransitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitionName

`func (o *WorkflowValidationErrorAllOf) SetTransitionName(v string)`

SetTransitionName sets TransitionName field to given value.

### HasTransitionName

`func (o *WorkflowValidationErrorAllOf) HasTransitionName() bool`

HasTransitionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


